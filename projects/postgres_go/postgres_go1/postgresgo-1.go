package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbMaxIdleConns = 4
	dbMaxConns = 40
	totalWorker = 100
	csvFile = "../../data-source/csv/majestic_million.csv"
	dsn = "postgres://postgres@localhost:5432/?sslmode=disable"
)

var dataHeader = make([]string, 0)

func main()  {

	start := time.Now()

	db, err := sql.Open("postgres", dsn)
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

	
	db.SetMaxOpenConns(dbMaxConns)
    db.SetMaxIdleConns(dbMaxIdleConns)


	csvReader, csvFile, err := openCsvFile()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{})
	wg := new(sync.WaitGroup)

	go dispatchWorkers(db, jobs, wg)
	readCsvFileLinePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func readCsvFileLinePerLineThenSendToWorker(csvReader *csv.Reader, jobs chan []interface{}, wg *sync.WaitGroup) {
	for {
		row, err := csvReader.Read()
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

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		wg.Add(1)
		jobs <- rowOrdered
	}
	close(jobs)
}

// SQLQueryDebugString formats an sql query inlining its arguments
// The purpose is debug only - do not send this to the database!
// Sending this to the DB is unsafe and un-performant.
func SQLQueryDebugString(query string, args ...interface{}) string {
	var buffer bytes.Buffer
	nArgs := len(args)
	// Break the string by question marks, iterate over its parts and for each 
	// question mark - append an argument and format the argument according to
	// it's type, taking into consideration NULL values and quoting strings.
	for i, part := range strings.Split(query, "?") {
		buffer.WriteString(part)
		if i < nArgs {
			switch a := args[i].(type) {
			case int64:
				buffer.WriteString(fmt.Sprintf("%d", a))
			case float64:
				buffer.WriteString(fmt.Sprintf("%v", a))
			case bool:
				buffer.WriteString(fmt.Sprintf("%t", a))
			case sql.NullBool:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%t", a.Bool))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullInt64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%d", a.Int64))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullString:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%q", a.String))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullFloat64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%v", a.Float64))
				} else {
					buffer.WriteString("NULL")
				}
			default:
				buffer.WriteString(fmt.Sprintf("%q", a))
			}
		}
	}
	return buffer.String()
}

func openCsvFile() (*csv.Reader, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}

func dispatchWorkers(db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex < totalWorker; workerIndex++ {
		go func (workerIndex int, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup)  {
			counter := 1
			for job := range jobs {
				doTheJob(workerIndex, counter, db, job)
				wg.Done()
				counter ++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func doTheJob(workerIndex, counter int, db *sql.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func ()  {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			conn, err := db.Conn(context.Background())
			if err != nil {
				log.Fatal(err.Error())
			}

			query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)", strings.Join(dataHeader, ","), strings.Join(generateDollarSigns(len(dataHeader)), ","),)
			// fmt.Println(SQLQueryDebugString(query, values...));
			_, err = conn.ExecContext(context.Background(), query, values...)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = conn.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}
	if counter%100 == 0 {
		log.Println("=> worker", workerIndex, "inserted", counter, "data")
	}
}

func generateDollarSigns(n int) []string {
	s := make([]string, 0)
	for i := 1; i < n+1; i++ {
		s = append(s, "$"+strconv.Itoa(i))
	}

	return s
}
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
	   old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
 }