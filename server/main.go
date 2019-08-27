package main

import (
	"fmt"
	"log"
	"myproto/model"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			defer c.Close()
			buffer := make([]byte, 1024)
			rn, err := c.Read(buffer)
			if err != nil {
				log.Println(err)
			}
			buffer = buffer[:rn]
			fmt.Println("Receive:", rn)

			person := &model.Person{}
			err = proto.Unmarshal(buffer, person)
			if err != nil {
				log.Println(err)
			}
			person.Active = true
			data, err := proto.Marshal(person)
			if err != nil {
				log.Println(err)
			}
			sn, err := c.Write(data)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Send:", sn)
		}(conn)
	}
}
