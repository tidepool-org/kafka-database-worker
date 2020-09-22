package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"github.com/segmentio/kafka-go"
	"encoding/json"

	"github.com/tidepool.org/kafka-database-worker/models"


	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var (
	DBContextTimeout = time.Duration(20)*time.Second
	KafkaContextTimeout = time.Duration(60)*time.Second

	Partition = 0
	HostStr, _ = os.LookupEnv("KAFKA_BROKERS")
	GroupId = "Tidepool-Mongo-Consumer29"
	//MaxMessages = 33100000
	MaxMessages = 100000
	WriteCount = 50000
	DeviceDataNumWorkers = 3
)




func init() {
	orm.SetTableNameInflector(func(s string) string {
		return  s
	})
}

func NewDbContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), DBContextTimeout)
	return ctx
}

func NewKafkaContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), KafkaContextTimeout)
	return ctx
}

type dbLogger struct { }

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	b, _ := q.FormattedQuery()
	s := string(b)
	if strings.Contains(s, "wizard") {
		fmt.Println("Query: ", s)
	}
	return nil
}


func connectToDatabase() *pg.DB {
	// Connect to db
	user, _ := os.LookupEnv("TIMESCALEDB_USER")
	password, _ := os.LookupEnv("TIMESCALEDB_PASSWORD")
	host, _ := os.LookupEnv("TIMESCALEDB_HOST")
	db_name, _ := os.LookupEnv("TIMESCALEDB_DBNAME")


	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=allow", user, password, host, db_name)
	opt, err := pg.ParseURL(url)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	fmt.Println("Trying to connect to db")

	ctx := NewDbContext()

	//db.AddQueryHook(dbLogger{})


	// Check if connection credentials are valid and PostgreSQL is up and running.
	if err := db.Ping(ctx); err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	fmt.Println("Connected successfully")

	return db
}


func worker(wg *sync.WaitGroup, db orm.DB, id int, jobs <-chan []interface{}, results chan<- bool) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", "len: ", cap(jobs))
		db.Insert(j)
		if err := db.Insert(j...); err != nil {
			// error has occurred
			fmt.Println("worker", id, "finished job - insert error", err)
			results <- true
		} else {
			results <- false
		}
	}
	fmt.Println("worker", id, "Completed");
	wg.Done()
}

func result(done chan bool, results <-chan  bool) {
	for result := range results {
		fmt.Println("Got Result for: ", result)
	}
	done <- true
}

func sendToDB(modelMap map[string][]interface{}, jobs chan <- []interface{}, count int,
	filtered int, decodingErrors int, deltaTime int64, topic string) {
	dataReceived := false
	recs := 0
	for _, val := range modelMap {
		if len(val) > 0 {
			jobs <- val
			dataReceived = true
		}
		recs += len(val)
	}
	//fmt.Printf("Delta Seconds:  kafak (ms): %d,  Timeseries (ms): %d\n",  kafkaDeltaTime/1000000, timeseriesDeltaTime/1000000)
	//fmt.Printf("Duration seconds: %f,  kafak (ms): %d,  Timeseries (ms): %d\n", time.Now().Sub(startTime).Seconds(), kafkaTime/1000000, timeseriesTime/1000000)
	if dataReceived {
		fmt.Printf("Received data:  %d\n", recs)
		fmt.Printf("Topic: %s, DeltaTime: %d,  Messages: %d,  filtered: %d,  decodingErrors: %d\n", topic, deltaTime/1000000, count+1, filtered, decodingErrors)
	} else {
		fmt.Printf("No data received\n")
		fmt.Printf("Topic: %s, DeltaTime: %d,  Messages: %d,  filtered: %d,  decodingErrors: %d\n", topic, deltaTime/1000000, count+1, filtered, decodingErrors)

	}

}

func createWorkers(numWorkers int, db orm.DB, jobs <- chan []interface{}, results chan <- bool) {
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		//fmt.Println("Created worker: ", i)
		go worker(&wg, db, i, jobs, results)
	}
	wg.Wait()
	close(results)
}

func readFromQueue(wg *sync.WaitGroup, db orm.DB, topic string, numWorkers int) {
	jobs := make(chan []interface{}, 5)
	results := make(chan bool)
	done := make(chan bool)

	fmt.Println("Reading topic: ", topic)


	go result(done, results)


	go createWorkers(numWorkers, db, jobs, results)

	//maxMessages :=  0
	prevTime := time.Now()
	//userFilters := map[string]bool {
	//	"c6505473f9": true,
	//	"9044a6953b": true,
	//	"298b233138": true,
	//}

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	groupid := GroupId + "." + topic
	fmt.Printf("Connecting to broker: %s,  topic: %s,  groupid: %s\n", HostStr, topic, groupid)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{HostStr},
		Topic:     topic,
		GroupID:   groupid,
		Partition: Partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		CommitInterval: 10*time.Second,
	})
	defer func() {
		if re := recover(); re != nil {
			fmt.Println("Recovered in read from queue", re)
		}
	}()


	modelMap := make(map[string][]interface{})
	filtered := 0
	decodingErrors := 0

	for i:=0; i<MaxMessages; i++ {
		m, err := r.ReadMessage(NewKafkaContext())
		if err != nil {
			fmt.Println(topic, "Timeout fetching message: \n", err)
			deltaTime := time.Now().Sub(prevTime).Nanoseconds()
			prevTime = time.Now()
			sendToDB(modelMap, jobs, i, filtered, decodingErrors, deltaTime, topic)
			modelMap = make(map[string][]interface{})
			continue
		}

		if (i+1) % WriteCount == 0 {
			deltaTime := time.Now().Sub(prevTime).Nanoseconds()
			prevTime = time.Now()
			sendToDB(modelMap, jobs, i, filtered, decodingErrors, deltaTime, topic)
			modelMap = make(map[string][]interface{})
		}
		var rec map[string]interface{}
		if err := json.Unmarshal(m.Value, &rec); err != nil {
			fmt.Println(topic, "Error Unmarshalling", err)
		} else {
			after_field, data_rec_ok := rec["after"]
			if data_rec_ok && after_field != nil {
			    var data map[string]interface{}
			    data_string := fmt.Sprintf("%v", after_field)
				if err := json.Unmarshal([]byte(data_string), &data); err != nil {
					//fmt.Println(topic, "Error Unmarshalling after field", err)
				} else {
					model, err := models.DecodeModel(data, topic)
					if err != nil {
						decodingErrors += 1
						//fmt.Println(topic, "Overall decoding error:", err)
					} else {
						if model != nil {
							_, ok := modelMap[model.GetType()]
							if !ok {
								modelMap[model.GetType()] = make([]interface{}, 0)
							}
							modelMap[model.GetType()] = append(modelMap[model.GetType()], model)
						} else {
							filtered += 1
						}
					}
				}
			}
		}
		//r.CommitMessages(context.Background(), m)
	}

	r.Close()
	wg.Done()
}

func main() {
	topics, _ := os.LookupEnv("KAFKA_TOPIC")

	fmt.Println("In main")
	time.Sleep(10 * time.Second)
	fmt.Println("Finished sleep")

	startTime := time.Now()
	db := connectToDatabase()
	defer db.Close()

	var wg sync.WaitGroup
	i := 1
	for _, topic := range strings.Split(topics, ",") {
		wg.Add(1)
		i++
		numWorkers := 1
		if strings.HasSuffix(topic, "Data") {
			numWorkers = DeviceDataNumWorkers
			go readFromQueue(&wg, db, topic, numWorkers)
		}

	}
	wg.Wait()

	fmt.Printf("Duration in seconds: %f\n", time.Now().Sub(startTime).Seconds())
	// Hack - do not quit for now
	fmt.Println("Sleeping until the end of time")
	for {
		time.Sleep(100 * time.Second)
	}
}