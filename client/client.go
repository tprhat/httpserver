package main

import (
	"log"
	"net"
)

func main() {
	Client()
}
func Client() {

	conn, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()
	data := []byte("Hello world")
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
