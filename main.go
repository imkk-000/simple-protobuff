package main

import (
	"fmt"
	"myproto/model"
	"strings"

	"github.com/golang/protobuf/proto"
)

// support only name and age
func apiV1() {
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
	fmt.Println("new.person.name:", newPerson.GetName())
	fmt.Println("new.person.age:", newPerson.GetAge())
}

// support name, age, and activities
func apiV2() {
	// create new protobuf struct
	person := &model.Person{
		Name: "imkk-000",
		Age:  111,
		Activities: []*model.Activity{
			{
				Name: "Workout",
				Time: 60,
			},
			{
				Name: "Sleep",
				Time: 60 * 8,
			},
		},
	}
	// try to get name, age, and activities from struct with protobuf function generator
	fmt.Println("person.name:", person.GetName())
	fmt.Println("person.age:", person.GetAge())
	fmt.Println("person.activities:", len(person.GetActivities()))
	for index, activity := range person.GetActivities() {
		fmt.Println("person.activities", index, ".name:", activity.GetName())
		fmt.Println("person.activities", index, ".time:", activity.GetTime())
	}
	// convert protobuf struct to binary and prepare for send to another service
	sendData, err := proto.Marshal(person)
	fmt.Println("sendData:", sendData)
	fmt.Println("sendData(string):", strings.TrimSpace(string(sendData)))
	fmt.Println("sendErr:", err)
	// convert sendData to receiveData
	newPerson := model.Person{}
	err = proto.Unmarshal(sendData, &newPerson)
	fmt.Println("receiveErr:", err)
	// show receive data name and age
	fmt.Println("new.person.name:", newPerson.GetName())
	fmt.Println("new.person.age:", newPerson.GetAge())
	fmt.Println("new.person.activities:", len(newPerson.GetActivities()))
	for index, activity := range newPerson.GetActivities() {
		fmt.Println("new.person.activities", index, ".name:", activity.GetName())
		fmt.Println("new.person.activities", index, ".time:", activity.GetTime())
	}
}

func main() {
	// welcome message
	fmt.Println("Hello Protobuf in Golang")
	apiV2()
}
