package main

import (
	"bufio"
	"github.com/segmentio/kafka-go"
	"log"

	"context"
	"encoding/json"
	"fmt"

	"github.com/tidepool.org/kafka-database-worker/models"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var (
	DBContextTimeout = time.Duration(20)*time.Second
	KafkaContextTimeout = time.Duration(60)*time.Second

	Partition = 0
	HostStr, _ = os.LookupEnv("KAFKA_BROKERS")
	GroupId = "Tidepool-Mongo-Consumer01"
	//MaxMessages = 33100000
	MaxMessages = 40000000
	WriteCount = 50000
	DeviceDataNumWorkers = 2
	Local = false
	StartOffset int64 = 0
	NumPartitions = 50
)

type UpdateRec struct {
	patch string
	id string
}



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
	if Local {
		url = fmt.Sprintf("postgres://postgres@localhost:5432/postgres?sslmode=disable")
	}
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


func processInsert(db orm.DB, models []interface{}) {
	fmt.Printf("started insert  len: %d \n", len(models))
	if err := db.Insert(models...); err != nil {
		// error has occurred
		fmt.Println("finished - insert error", err)
	}
}

func processUpdates(db orm.DB, updates []UpdateRec) error {
	models.GetModelTypes()
	fmt.Printf("started updates  len: %d \n", len(updates))
	for _, update := range updates {

		for _, modelType := range models.GetModelTypes() {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(update.patch), &data); err != nil {
				//fmt.Println(topic, "Error Unmarshalling after field", err)
				continue
			}

			model, metadata, err := models.DecodeDeviceModelWithType(data["$set"], modelType.Type)
			if err != nil {
				continue
			}

			res, err := db.Model(model).
				Column(convertKeys(metadata.Keys, modelType.TagMap)...).
				Where("internal_mongo_id = ?", update.id).
				Update()
			if err != nil {
				continue
			}
			if res.RowsAffected() == 0 {
				continue
			}
			break
		}
	}
	return nil
}

func convertKeys(keys []string, tagMap map[string]string) []string {
	var convertedKeys []string
	for _, key := range keys {
		tag, ok := tagMap[key]
		if ok != false {
			convertedKeys = append(convertedKeys, tag)
		}
	}
	return convertedKeys
}

func processDeletes(db orm.DB, deletes []string) error {
	// Have to go through all the tables and see if deletes work
	fmt.Printf("started deletes  len: %d \n", len(deletes))
	ids := pg.Strings(deletes)
	numDeleted := 0
	for _, modelDef := range models.GetModels() {
		res, err := db.Model(modelDef).Where("internal_mongo_id IN (?)", ids).Delete()
        if err != nil {
        	return err
		}
		numDeleted += res.RowsAffected()
		if numDeleted >= len(deletes) {
			return nil
		}
	}

	fmt.Printf("Only Deleted %d of %d records", numDeleted, len(deletes))
	return errors.New("Did not delete all records")

}

func sendToDB(db orm.DB, modelMap map[string][]interface{}, updates []UpdateRec, deletes []string, count int,
	filtered int, decodingErrors int, deltaTime int64, topic string,
	insertsCount int, updatesCount int, deletesCount int) {
	recs := 0

	// First do inserts
	for _, val := range modelMap {
		if len(val) > 0 {
			processInsert(db, val)
			//fmt.Printf("Placed on jobs len: %d\n", len(val))
		}
		recs += len(val)
	}

	// Next do updates
	if len(updates) > 0 {
		if err := processUpdates(db, updates); err != nil {
			fmt.Println(err)
		}
	}

	// Finally - do deletes
	if len(deletes) > 0 {
		if err := processDeletes(db, deletes); err != nil {
			fmt.Println(err)
		}

	}

	if recs + len(updates) + len(deletes) > 0 {
		fmt.Printf("New inserts:  %d, updates: %d,  deletes: %d\n", recs, len(updates), len(deletes))
	} else {
		fmt.Printf("No data received\n")
	}
	fmt.Printf("Topic: %s, DeltaTime: %d,  Messages: %d,  filtered: %d,  decodingErrors: %d\n", topic, deltaTime/1000000, count+1, filtered, decodingErrors)
	fmt.Printf("Totals:  Inserts: %d  Updates: %d  Deletes: %d\n", insertsCount, updatesCount, deletesCount)
	fmt.Println("")
}

func readFromQueue(wg *sync.WaitGroup, db orm.DB, topic string, numWorkers int) {

	fmt.Println("Reading topic: ", topic)


	prevTime := time.Now()

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	var r *kafka.Reader
	var scanner *bufio.Scanner


	brokers := []string{HostStr}
	//topics := []string{topic}
	groupid := GroupId + "." + topic

	if !Local {

		// If we do not start at zero - commit offsets at new location
		// This will not work for multiple replicas
		//if StartOffset != 0 {
		//	ConsumerGroupOverwriteOffsets(StartOffset, NumPartitions, groupid, brokers, topics)
		//}

		// Connect to broker
		fmt.Printf("Connecting to broker: %s,  topic: %s,  groupid: %s\n", HostStr, topic, groupid)
		r = kafka.NewReader(kafka.ReaderConfig{
			Brokers:        brokers,
			Topic:          topic,
			GroupID:        groupid,
			//Partition:      Partition,
			MinBytes:       10e3, // 10KB
			MaxBytes:       10e6, // 10MB
			CommitInterval: 10 * time.Second,
		})
		defer func() {
			if re := recover(); re != nil {
				fmt.Println("Recovered in read from queue", re)
			}
		}()

	} else {

		filename := "ll5"
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)

	}


	modelMap := make(map[string][]interface{})
	var updates []UpdateRec
	var deletes []string
	filtered := 0
	decodingErrors := 0
	insertsCount := 0
	deletesCount := 0
	updatesCount := 0
	messageCount := 0

	for messageCount =0; messageCount <MaxMessages; messageCount++ {

		var b []byte
		if !Local {

			m, err := r.ReadMessage(NewKafkaContext())
			if err != nil {
				fmt.Println(topic, "Timeout fetching message: \n", err)
				deltaTime := time.Now().Sub(prevTime).Nanoseconds()
				prevTime = time.Now()
				sendToDB(db, modelMap, updates, deletes, messageCount, filtered, decodingErrors, deltaTime, topic, insertsCount, updatesCount, deletesCount)
				modelMap = make(map[string][]interface{})
				deletes = []string{}
				updates = []UpdateRec{}
				continue
			}
			b = m.Value

		} else {
			e  := scanner.Scan()
			if e != true {
				fmt.Println("Scanner done")
				break
			}
			line := scanner.Text()
			b = []byte(line)
		}

		if (messageCount+1) % WriteCount == 0 {
			deltaTime := time.Now().Sub(prevTime).Nanoseconds()
			prevTime = time.Now()
			sendToDB(db, modelMap, updates, deletes, messageCount, filtered, decodingErrors, deltaTime, topic, insertsCount, updatesCount, deletesCount)
			modelMap = make(map[string][]interface{})
			deletes = []string{}
			updates = []UpdateRec{}
		}

		// Rec contains the entire kafka record
		var rec map[string]interface{}
		if err := json.Unmarshal(b, &rec); err != nil {
			fmt.Println(topic, "Error Unmarshalling", err)
		} else {

			// This is a Create Event
			after_field, data_field_ok := rec["after"]
			patch_field, patch_field_ok := rec["patch"]
			filter_field, filter_field_ok := rec["filter"]
			if data_field_ok && after_field != nil {

				// Data contains the contents of the after field which is a json object
			    var data map[string]interface{}
			    after_field_string := fmt.Sprintf("%v", after_field)
				if err := json.Unmarshal([]byte(after_field_string), &data); err != nil {
					//fmt.Println(topic, "Error Unmarshalling after field", err)
				} else {

					// Get the internal mongoId and then decode the object according to our models
					model, _, err := models.DecodeModel(data, topic)
					// Do some checks to see if we can fix potential errors
					if err != nil {
						if strings.Contains(err.Error(), "conversionOff") {
							v := reflect.ValueOf(data["conversionOffset"])
							if v.Kind() == reflect.Map {
								// get the value that the pointer v points to.
								item := v.MapIndex(reflect.ValueOf("$numberLong"))
								if item.IsValid() {
									conversionOffset, e := strconv.Atoi(fmt.Sprintf("%v", item))
									if e == nil {
										data["conversionOffset"] = conversionOffset
										model, _, err = models.DecodeModel(data, topic)
									}
								}
							}
						}
					}
					if err != nil {
						decodingErrors += 1
						fmt.Println(topic, "Overall decoding error:", err)
					} else {
						if model != nil {
							_, ok := modelMap[model.GetType()]
							if !ok {
								modelMap[model.GetType()] = make([]interface{}, 0)
							}
							modelMap[model.GetType()] = append(modelMap[model.GetType()], model)
							insertsCount += 1
						} else {
							filtered += 1
						}
					}
				}

			// Check for update Event
			} else if patch_field_ok && patch_field != nil && filter_field_ok && filter_field != nil {

				if id := models.GetMongoIdFromFilterField(filter_field); id != nil {
					updates = append(updates, UpdateRec{patch: patch_field.(string), id: *id})
					updatesCount += 1
				}

			// Check for delete Event
			} else if filter_field_ok && filter_field != nil {

				if id := models.GetMongoIdFromFilterField(filter_field); id != nil {
					deletes = append(deletes, *id)
					deletesCount += 1
				}

			// Not a valid kafka event
			} else {
				fmt.Println("Illegal Kafka record.  After, patch and filter fields are nil")
			}
		}

	}
	fmt.Println(topic, "Finishing processing messages - cleanup")
	deltaTime := time.Now().Sub(prevTime).Nanoseconds()
	prevTime = time.Now()
	sendToDB(db, modelMap, updates, deletes, messageCount, filtered, decodingErrors, deltaTime, topic, insertsCount, updatesCount, deletesCount)

	if !Local {
		r.Close()
	}
	wg.Done()
}

func ConsumerGroupOverwriteOffsets(offset int64, numPartitions int, consumerGroupID string, brokers []string, topics []string) {
	group, err := kafka.NewConsumerGroup(kafka.ConsumerGroupConfig{
		ID:      consumerGroupID,
		Brokers: brokers,
		Topics:  topics,
	})
	if err != nil {
		fmt.Printf("error creating consumer group %s: %+v\n", consumerGroupID, err)
		os.Exit(1)
	}
	defer group.Close()

	gen, err := group.Next(context.TODO())
	if err != nil {
		fmt.Printf("error getting next generation: %+v\n", err)
		os.Exit(1)
	}
	var offsets map[string]map[int]int64
	offsets = make(map[string]map[int]int64, numPartitions)
	for i := 0; i < numPartitions; i++ {
		offsets[consumerGroupID][i] = offset
	}
	err = gen.CommitOffsets(offsets)
	if err != nil {
		fmt.Printf("error committing offsets next generation: %+v\n", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("In main")
	topics, _ := os.LookupEnv("KAFKA_TOPIC")
	if Local {
		topics = "deviceData"
	} else {
		time.Sleep(10 * time.Second)
	}

	fmt.Println("Finished sleep")

	startTime := time.Now()
	db := connectToDatabase()
	defer db.Close()

	var wg sync.WaitGroup
	i := 1
	for _, topic := range strings.Split(topics, ",") {
		if strings.HasSuffix(topic, "Data") {
			wg.Add(1)
			i++
			numWorkers := 1
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