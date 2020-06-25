package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/segmentio/kafka-go"
	"encoding/json"

	"github.com/tidepool.org/kafka-database-worker/models"


	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var (
	ContextTimeout = time.Duration(20)*time.Second
)




func init() {
	orm.SetTableNameInflector(func(s string) string {
		return  s
	})
}

func NewDbContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), ContextTimeout)
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


func writeToDatabase() {
	// Connect to db
	user := "postgres"
	password, _ := os.LookupEnv("TIMESCALEDB_PASSWORD")
	host := "timescaledb-single.timescaledb.svc.cluster.local"
	db_name := "postgres"


	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=allow", user, password, host, db_name)
	opt, err := pg.ParseURL(url)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	defer db.Close()
	fmt.Println("Trying to connect to db")

	ctx := NewDbContext()

	//db.AddQueryHook(dbLogger{})


	// Check if connection credentials are valid and PostgreSQL is up and running.
	if err := db.Ping(ctx); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Connected successfully")

	readFromQueue(db)
}

func readFromQueue(db orm.DB) {
	topic := "database"
	partition := 0
	host := "kafka-kafka-bootstrap.kafka.svc.cluster.local"
	port := 9092
	hostStr := fmt.Sprintf("%s:%d", host,port)
	maxMessages := 500000
	startTime := time.Now()

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{hostStr},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		CommitInterval: 10*time.Second,
	})



	kafkaTime := int64(0)
	timeseriesTime := int64(0)
	var modelArray []interface{}
	for i:=0; i<maxMessages; i++ {
		archived := 0
		insertErrors := 1
		kafkaStartTime := time.Now()
		m, err := r.ReadMessage(context.Background())
		kafkaTime += time.Now().Sub(kafkaStartTime).Nanoseconds()
		if err != nil {
			fmt.Println("Error fetching message: ", err)
			break
		}

		if (i-1) % 1000 == 0 {
			timeseriesStartTime := time.Now()
			if _, err := db.Model(modelArray...).Insert(); err != nil {
				fmt.Println("Error inserting: ", err)
				insertErrors += 1
			}
			timeseriesTime += time.Now().Sub(timeseriesStartTime).Nanoseconds()
			fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
			fmt.Printf("Duration seconds: %f,  kafakTime (ms): %f,  TimeseriesTime (ms): %f\n", time.Now().Sub(startTime).Seconds(), kafkaTime/1000000, timeseriesTime/1000000)
			fmt.Printf("Messages: %d,  Archived: %d, insertErrors: %d", i, archived, insertErrors)
		}
		var rec map[string]interface{}
		if err := json.Unmarshal(m.Value, &rec); err != nil {
			fmt.Println("Error Unmarshalling", err)
		} else {
			source, source_ok := rec["source"]
			data, data_ok := rec["data"]
			if data_ok && source_ok && source == "database"{

				if model := models.DecodeModel(data); model != nil {
					modelArray = append(modelArray, model)
				} else {
					archived += 1;
				}
			}
		}
		//r.CommitMessages(context.Background(), m)
	}

	r.Close()
}

func main() {
	fmt.Println("In main")
	time.Sleep(10 * time.Second)
	fmt.Println("Finished sleep")

	startTime := time.Now()
	writeToDatabase()
	fmt.Printf("Duration in seconds: %f\n", time.Now().Sub(startTime).Seconds())
	// Hack - do not quit for now
	fmt.Println("Sleeping until the end of time")
	for {
		time.Sleep(10 * time.Second)
	}
}