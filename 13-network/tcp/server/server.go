package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("Error listening on 4000 : %v\n", err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("A new connection is established")
		if err != nil {
			log.Fatalf("Acceping connections : %v\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buff := make([]byte, 512)
		_, err := conn.Read(buff)
		if err != nil {
			log.Fatalf("Error reading data : %v\n", err)
		}
		fmt.Printf("Received data : %v\n", string(buff))
		//echoing the message to the client
		conn.Write(buff)
	}
}
