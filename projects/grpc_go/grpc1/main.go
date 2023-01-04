package main

import (
	"addressbook/protos"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)


func main()  {
	p := protos.Person{
		Id: 1234,
		Name: "John Doe",
		Email: "jdoe@example.com",
		Phones: []*protos.Person_PhoneNumber{
			{Number: "555-4321", Type: protos.Person_HOME},
		},
	}

	book := &protos.AddressBook{}
	book.People = append(book.People, &p)

	// write the new address book back to disk
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile("addressbook", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// read the existing address book.
	in, err := ioutil.ReadFile("addressbook")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	addresslist := &protos.AddressBook{}
	if err := proto.Unmarshal(in, addresslist); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(addresslist)
}