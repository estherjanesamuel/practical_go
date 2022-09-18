package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type person struct {
    ID int
    FirstName, LastName string
    Age uint8
}

func main() {

    file, err := os.Open("csv1.csv")
    if err != nil {
        log.Panicln(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    var people []person
    records, _ := reader.ReadAll()
    for _, record := range records[1:] {
        id, _ := strconv.Atoi(record[0])
        age, _ := strconv.Atoi(record[3])
        person := person {
            ID: id,
            FirstName: record[1],
            LastName: record[2],
            Age: uint8(age),
        }
        people = append(people, person)
    }
    if len(people) > 4 {
        fmt.Println(people)
    }
}


