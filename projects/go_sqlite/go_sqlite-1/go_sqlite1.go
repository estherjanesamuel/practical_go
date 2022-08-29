package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/dixonwille/wmenu/v5"
)
func main()  {

	db, err := sql.Open("sqlite3", "./names.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println()
		menu := wmenu.NewMenu("What would you like to do?")
		menu.Action(func(o []wmenu.Opt) error {handleFunc(db, o); return nil})
		
		menu.Option("Add a new Person", 0, true, nil)
		menu.Option("Get a Person", 1, false, nil)
		menu.Option("Edit a Person", 2, false, nil)
		menu.Option("Delete a Person", 3, false, nil)
		menu.Option("Exit", 4, false, nil)
		err = menu.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func handleFunc(db *sql.DB,o []wmenu.Opt) {
	switch o[0].Value {
	case 0:
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter a first name: ")
		fn, _ := reader.ReadString('\n')
		fn = strings.TrimSuffix(fn,"\n")

		fmt.Println("Enter a last name: ")
		ln, _ := reader.ReadString('\n')
		ln = strings.TrimSuffix(ln,"\n")

		fmt.Println("Enter a email address: ")
		e, _ := reader.ReadString('\n')
		e = strings.TrimSuffix(e,"\n")

		fmt.Println("Enter a ip address: ")
		ip, _ := reader.ReadString('\n')
		ip = strings.TrimSuffix(ip,"\n")


		person := person{
			id: 0,
			first_name: fn,
			last_name: ln,
			email: e,
			ip_address: ip,
		}

		add(db, person)
	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a name to search for : ")
		searchString, _ := reader.ReadString('\n')
    	searchString = strings.TrimSuffix(searchString, "\n")
		people := search(db, searchString)

		fmt.Printf("Found %v results\n", len(people))

		for _, ourPerson := range people {
			fmt.Printf("\n----\nID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s\nIP Address: %s\n",ourPerson.id, ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address)
		}
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an id to edit: ")
		id, _ := reader.ReadString('\n')
		currentPerson := getById(db, id)

		fmt.Printf("first name (currently `%s`): ", currentPerson.first_name)
		fn, _ := reader.ReadString('\n')
		if fn != "\n" {
			currentPerson.first_name = strings.TrimSuffix(fn,"\n")
		}
		
		fmt.Printf("last name (currently `%s`): ", currentPerson.last_name)
		ln, _ := reader.ReadString('\n')
		if ln != "\n" {
			currentPerson.last_name = strings.TrimSuffix(ln,"\n")
		}

		fmt.Printf("email (currently `%s`): ", currentPerson.email)
		e, _ := reader.ReadString('\n')
		if e != "\n" {
			currentPerson.email = strings.TrimSuffix(e,"\n")
		}

		fmt.Printf("ip address (currently `%s`): ", currentPerson.ip_address)
		ip, _ := reader.ReadString('\n')
		if ip != "\n" {
			currentPerson.ip_address = strings.TrimSuffix(ip,"\n")
		}

		affected := update(db, currentPerson)

		if affected == 1 {
			fmt.Println("one row affected (updated)")
		}
	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an id to delete: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSuffix(id, "\n")

		affected := deleteById(db, id)
		if affected == 1 {
			fmt.Println("one row affected (deleted)")
		}
	case 4:
		fmt.Println("Good Bye!")
		os.Exit(1)
	}
}

func deleteById(db *sql.DB, id string) int64{
	stmt, _ := db.Prepare("DELETE FROM people WHERE id=?")
	res, err := stmt.Exec(id) 
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return affected
}

func update(db *sql.DB, person person) int {
	stmt, _ := db.Prepare("UPDATE people SET first_name=?, last_name=?, email=?, ip_address=? WHERE id = ?;	")
	res, err := stmt.Exec(person.first_name, person.last_name, person.email, person.ip_address, person.id) 
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(affected)
}

func getById(db *sql.DB, id string) person {
	rows, err := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE id ='" + id + "'")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	person := person{}
	for rows.Next() {
		rows.Scan(&person.id,&person.first_name,&person.last_name, &person.email, &person.ip_address)
	}
	return person
}

func search(db *sql.DB, searchString string) []person {
	
	rows, err := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE first_name like '%" + searchString + "%' OR last_name like '%" + searchString + "%'")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]person, 0)

	for rows.Next() {
		ourPerson := person{}
		err = rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people
}

func add(db *sql.DB, person person) {
	stmt, _ := db.Prepare("INSERT INTO people (id,first_name, last_name, email, ip_address) VALUES(?, ?, ?, ?, ?)")
	stmt.Exec(nil, person.first_name, person.last_name, person.email, person.ip_address) 
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", person.first_name, person.last_name)
}

type person struct {
	first_name, last_name, email, ip_address string
	id int
}
