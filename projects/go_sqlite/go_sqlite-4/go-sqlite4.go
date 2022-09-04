package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbMaxIdleConns = 1
	dbMaxConns = 1
	totalWorker = 100
	csvFile = "../../data-source/csv/majestic_million.csv"
	// dsn = "file:./workshop.db?cache=shared&_journal_mode=WAL&auto_vacuum=0&synchronous=OFF"
	dsn = "file:./poc.db"
)

var dataHeader = make([]string, 0)

func main()  {

	start := time.Now()

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS domain (GlobalRank int,TldRank int, Domain varchar(255), TLD varchar(255), RefSubNets int, RefIPs int, IDN_Domain varchar(255), IDN_TLD varchar(255), PrevGlobalRank int, PrevTldRank int, PrevRefSubNets int, PrevRefIPs int )")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("PRAGMA journal_mode = OFF;")
	db.Exec("PRAGMA synchronous = 0;")
	db.Exec("PRAGMA cache_size = 1000000;")
	db.Exec("PRAGMA locking_mode = EXCLUSIVE;")
	db.Exec("PRAGMA temp_store = MEMORY;")
	
	
	db.SetMaxOpenConns(dbMaxConns)
    db.SetMaxIdleConns(dbMaxIdleConns)

	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	_, err = r.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(dataHeader) == 0 {
			dataHeader = row
			continue
		}
	}
	


	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

