package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
	//"github.com/go-pg/pg/v10/orm"
)

var (
	ContextTimeout = time.Duration(20)*time.Second
)

type Basal struct {
	time              time.Time  `sql:"type:timestamptz"`

	uploadId          string

	deliveryType      string
	duration          int64
	expectedDuration  int64
	rate              float64
	percent           float64
	scheduleName      string

}
func NewDbContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), ContextTimeout)
	return ctx
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

	// Check if connection credentials are valid and PostgreSQL is up and running.
	if err := db.Ping(ctx); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Connected successfully")

	basal := &Basal{
		time: time.Now(),
		uploadId: "upid4545",
		deliveryType: "automated",
		duration: 50,
		rate: 45.45,
	}

	fmt.Println("Inserting into db")
	err = db.Insert(basal)
	if err != nil {
		fmt.Println("Error inserting: ", err)
		return
	}
	fmt.Println("inserted successfully")
}

func main() {
	fmt.Println("In main")
	time.Sleep(10 * time.Second)
	fmt.Println("Finished sleep")
	writeToDatabase()
	// Hack - do not quit for now
	fmt.Println("Sleeping until the end of time")
	for {
		time.Sleep(10 * time.Second)
	}
}