package main

import (
	"log"
	"myproto/model"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {
	// create connection
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to:", conn.LocalAddr().String())
	// create person data
	person := &model.Person{
		Name: "imkk-000",
		Age:  111,
		Activities: []*model.Activity{
			{
				Name: "Workout",
				Time: 60,
			},
			{
				Name: "Slepp",
				Time: 60 * 8,
			},
		},
	}
	log.Println("Active:", person.GetActive())
	// send person data to server
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	sn, err := conn.Write(data)
	if err != nil {
		log.Println(err)
	}
	log.Println("Sent:", sn, data)
	// receive person data from server
	buffer := make([]byte, 1024)
	rn, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
	}
	buffer = buffer[:rn]
	log.Println("Receive:", rn)
	newPerson := &model.Person{}
	err = proto.Unmarshal(buffer, newPerson)
	if err != nil {
		log.Println(err)
	}
	log.Println("Active:", newPerson.GetActive())
}
