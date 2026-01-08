package main

import (
	"log"
)

func main() {
	server := NewServer()

	log.Println("API running on :5000")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
