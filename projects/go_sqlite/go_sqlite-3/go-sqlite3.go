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

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbMaxIdleConns = 1
	dbMaxConns = 1
	totalWorker = 100
	csvFile = "../../data-source/csv/majestic_million.csv"
	dsn = "file:./poc.db"
	min_batch_size = 50
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
	
	/*
	
	db.Exec("PRAGMA journal_mode = WAL")
    db.Exec("PRAGMA synchronous = 0")
    db.Exec("PRAGMA cache_size = 1000000")
    db.Exec("PRAGMA locking_mode = EXCLUSIVE")
    db.Exec("PRAGMA temp_store = MEMORY")
	
	8695  insert / second
	
	// Use WAL : https://stackoverflow.com/a/35805826
	// db.Exec("PRAGMA journal_mode = WAL")
	
	// https://stackoverflow.com/questions/1711631/improve-insert-per-second-performance-of-sqlite
	db.Exec("PRAGMA auto_vacuum = 0")
	db.Exec("PRAGMA synchronous = OFF")
	
	*/
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
				// doTheJob(workerIndex, counter, db, job)
				JobStore(workerIndex, counter, db, job)
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

			query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)", strings.Join(dataHeader, ","), strings.Join(generateQuestionsMark(len(dataHeader)), ","),)
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

func TheJob(workerIndex, counter int, db *sql.DB, values []interface{}) {
	jobs := make([][]interface{}, 0)
	valueStr := fmt.Sprintf("(%s),", strings.Join(generateQuestionsMark(len(dataHeader)), ","))
	valueStr = strings.Repeat(valueStr,min_batch_size)
	valueStr = strings.TrimSuffix(valueStr, ",")
	valueArgs := make([]interface{}, 0)
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

			// query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)", strings.Join(dataHeader, ","), strings.Join(generateQuestionsMark(len(dataHeader)), ","),)
			// fmt.Println(SQLQueryDebugString(query, values...));
			// _, err = conn.ExecContext(context.Background(), query, values...)

			jobs = append(jobs, values)
			counter ++
			if len(jobs) == min_batch_size {
				for _, job := range jobs {
					valueArgs = append(valueArgs, job)
				}
				stmt := fmt.Sprintf("INSERT INTO domain (%s) VALUES %s", strings.Join(dataHeader, ","), valueStr,)
				_, err = conn.ExecContext(context.Background(), stmt, valueArgs...)
				fmt.Println(stmt, valueArgs)
				
				jobs = nil
				
			}
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

func JobStore(workerIndex, counter int, db *sql.DB, values []interface{}) {
	domains := make([]domain, 0)
	vals := []interface{}{}
	valueStr := fmt.Sprintf("(%s),", strings.Join(generateQuestionsMark(len(dataHeader)), ","))
	valueStr = strings.Repeat(valueStr,min_batch_size)
	valueStr = strings.TrimSuffix(valueStr, ",")

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

			d := domain{
				GlobalRank: values[0].(string),
				TldRank : values[1].(string),
				Domain : values[2].(string),
				TLD : values[3].(string),
				RefSubNets : values[4].(string),
				RefIPs : values[5].(string),
				IDN_Domain : values[6].(string),
				IDN_TLD : values[7].(string),
				PrevGlobalRank : values[8].(string),
				PrevTldRank : values[9].(string),
				PrevRefSubNets : values[10].(string),
				PrevRefIPs : values[11].(string),
				} 
			domains = append(domains, d)	
			counter ++

			if counter % 50 == 0 {
				for _, row := range domains {
					vals = append(vals, row.GlobalRank, row.TldRank, row.Domain, row.TLD, row.RefSubNets, row.RefIPs, row.IDN_Domain, row.IDN_TLD, row.PrevGlobalRank, row.PrevTldRank, row.PrevRefSubNets, row.PrevRefIPs )
				}
				fmt.Println(len(domains))
				// sqlStr := `INSERT INTO domain (%s) VALUES %s`
				// sqlStr = ReplaceSQL(sqlStr, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", len(domains))
				// sqlStr := fmt.Sprintf("INSERT INTO domain (%s) VALUES %s", strings.Join(dataHeader, ","), valueStr,)
			
				// //Prepare and execute the statement
				// stmt, _ := db.Prepare(sqlStr)
				// _, err := stmt.Exec(vals...)
				// if err != nil {
				// 	log.Fatal(err)	
				// }
			}
			// query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)", strings.Join(dataHeader, ","), strings.Join(generateQuestionsMark(len(dataHeader)), ","),)
			// // vals = append(vals,values...)

			// if counter % 1000 == 0 {
			// 	// fmt.Println(SQLQueryDebugString(query, domains),)
				
			// }

			// if counter % 10000 == 0 {
			// 	fmt.Println(SQLQueryDebugString(query, values...),)
				
			// }
			//_, err = conn.ExecContext(context.Background(), query, values...)
			/*
			vals := []interface{}{}
for _, row := range data {
    vals = append(vals, row.FirstName, row.LastName, row.Age)
}

sqlStr := `INSERT INTO test(column1, column2, column3) VALUES %s`
sqlStr = ReplaceSQL(sqlStr, "(?, ?, ?)", len(data))

//Prepare and execute the statement
stmt, _ := db.Prepare(sqlStr)
res, _ := stmt.Exec(vals...)
			*/

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

func ReplaceSQL(stmt, pattern string, len int) string {
    pattern += ","
    stmt = fmt.Sprintf(stmt, strings.Repeat(pattern, len))
    n := 0
    for strings.IndexByte(stmt, '?') != -1 {
        n++
        param := "$" + strconv.Itoa(n)
        stmt = strings.Replace(stmt, "?", param, 1)
    }
    return strings.TrimSuffix(stmt, ",")
}

func generateQuestionsMark(n int) []string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, "?")
	}

	return s
}

func runInsert(workerIndex int, db *sql.DB, values []interface{})  {
	
	count := len(values)
	// counts := 0
	// counts = counts + len(values)
	// fmt.Println("len values",len(values))
	// fmt.Println("total",counts)
	if count < min_batch_size {
		log.Fatal("count can't be less than min batch size")
	}
	valueStr := fmt.Sprintf("(%s),", strings.Join(generateQuestionsMark(len(dataHeader)), ","))
	valueStr = strings.Repeat(valueStr,min_batch_size)
	valueStr = strings.TrimSuffix(valueStr, ",")
	valueArgs := make([]interface{}, 0 , len(values))
	counter := 1

	for {
		var outerError error
		func(outerError *error) {
			defer func ()  {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()
			
			conn, err := db.Conn(context.Background())
			stmt := fmt.Sprintf("INSERT INTO domain (%s) VALUES %s", strings.Join(dataHeader, ","), valueStr,)
			if err != nil {
				log.Fatal(err.Error())
			}
			for i := 0; i < count / min_batch_size; i++ {
				for j := 0; j < min_batch_size; j++ {
					valueArgs = append(valueArgs, values[j])
					counter ++
				}
				fmt.Println(stmt, valueArgs)
				_, err = conn.ExecContext(context.Background(), stmt, valueArgs...)
			}



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

type domain struct {
	GlobalRank string
    TldRank string
    Domain string
    TLD string
    RefSubNets string
    RefIPs string
    IDN_Domain string
    IDN_TLD string
    PrevGlobalRank string
    PrevTldRank string
    PrevRefSubNets string
    PrevRefIPs string
}