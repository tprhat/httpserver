package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type player struct {
	name int
	hand hand
}

func main() {
	deck := buildDeck()
	// log.Print(deck)
	log.Print(len(deck))
	counter := 0
	players := make([]player, 2)
	ln, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Server started: %v\n", ln)

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if counter >= 2 {
			log.Printf("Rejecting connection from %s", conn.RemoteAddr())
			conn.Close()
			continue
		}
		log.Printf("Client %s connected!", conn.RemoteAddr())
		players[counter] = player{name: counter}
		if err != nil {
			log.Printf("Error accepting: %v", err)
			continue
		}
		if conn == nil {
			log.Print("nil")
		} else {
			go HandleClient(conn, &players[counter], &counter)
			counter++
			log.Printf("num of players = %d", counter)
			log.Printf("current players are: %v", players)
		}
	}
}

func HandleClient(conn net.Conn, player *player, counter *int) {
	defer conn.Close()
	log.Printf("player %d joined:", player.name)
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			*counter--
			if err == io.EOF {
				log.Printf("connection %d closed", player.name)
			} else {
				log.Printf("Error: client closed connection. Details: %v", err)
			}
			return
		}
		log.Printf("From: player #%d\tData: %s\n", player.name, buffer[:n])
		// _, err = conn.Write([]byte("Here is res:"))
		// if err != nil {
		// log.Print("Error sending response")
		// }
	}
}
