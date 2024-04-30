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
	fmt.Printf("Server started: %v\n", ln)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %v", err)
		}
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	deck := buildDeck()
	for _, card := range deck {
		log.Printf("%d: %s", card.num, card.color)
	}
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
		fmt.Printf("Data received: %s\n", buffer[:n])
	}
}
