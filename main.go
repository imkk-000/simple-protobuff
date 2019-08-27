package main

import (
	"fmt"
	"myproto/model"
	"strings"

	"github.com/golang/protobuf/proto"
)

func main() {
	// welcome message
	fmt.Println("Hello Protobuf in Golang")
	// create new protobuf struct
	person := model.Person{
		Name: "imkk-000",
		Age:  111,
	}
	// try to get name and age from struct with protobuf function generator
	fmt.Println("person.name:", person.GetName())
	fmt.Println("person.age:", person.GetAge())
	// convert protobuf struct to binary and prepare for send to another service
	sendData, err := proto.Marshal(&person)
	fmt.Println("sendData:", sendData)
	fmt.Println("sendData(string):", strings.TrimSpace(string(sendData)))
	fmt.Println("sendErr:", err)
	// convert sendData to receiveData
	newPerson := model.Person{}
	err = proto.Unmarshal(sendData, &newPerson)
	fmt.Println("receiveErr:", err)
	// show receive data name and age
	fmt.Println("new.person.name:", person.GetName())
	fmt.Println("new.person.age:", person.GetAge())
}
