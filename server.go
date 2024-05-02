package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type player struct {
	name int
	addr net.Addr
}

func main() {
	counter := 0
	players := make([]player, 2)
	ln, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Server started: %v\n", ln)
	for {
		conn, err := ln.Accept()
		log.Printf("Client %s connected!", conn.RemoteAddr())
		players[counter] = player{name: counter, addr: conn.RemoteAddr()}
		counter++
		if err != nil {
			log.Fatalf("Error accepting: %v", err)
		}
		log.Printf("num of players = %d", counter)
		log.Printf("current players are: %v", players)	
		go HandleClient(conn, conn.RemoteAddr())
	}
}

func HandleClient(conn net.Conn, remoteAddr net.Addr) {
	// deck := buildDeck()
	// for _, card := range deck {
	// 	log.Printf("%d: %s", card.num, card.color)
	// }
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("connection %s closed", remoteAddr)
			} else {
				log.Printf("Error: %s", err)
			}
			return
		}
		log.Printf("From: %s\tData: %s\n", remoteAddr, buffer[:n])
		_, err = conn.Write([]byte("Here is res:"))
		if err != nil {
			log.Print("Error sending response")
		}
	}
}
