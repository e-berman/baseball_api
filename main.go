package main

import (
	"log"
)

func main() {
	dbpool, err := NewDBPool()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.db.Close()

	if err := dbpool.InitializePlayerTable(); err != nil {
		log.Fatal(err)
	}

	server := NewServer(":4242", dbpool)
	server.StartServer()
}
