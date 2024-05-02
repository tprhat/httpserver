package main

import (
	"bufio"
	"log"
	"net"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := conn.Write([]byte(text + "\n)"))
		if text == ":wqa" {
			return
		}
		if err != nil {
			log.Printf("Error sending text")
			return
		}
	}
	// Write to server
	data := []byte("Hello world")
	_, err = conn.Write(data)
	if err != nil {
		log.Printf("Error writing data to the server: %v", err)
	}
	// Read from server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Response: %s", buffer[:n])
}
