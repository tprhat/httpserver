package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Server started: %v", ln)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Error: %v", err)
			return
		}
		fmt.Printf("Data received: %s\n", buffer[:n])
	}
}