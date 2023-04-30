package main

import (
	"fmt"
	"log"
	"strings"
)

var Fp string

func main() {
	var choice string

	fmt.Println("Do you want to import a csv file?: ")
	fmt.Println("[Y]: Yes")
	fmt.Println("[N]: No")
	fmt.Scanln(&choice)
	fmt.Println("You entered: ", choice)

	if strings.ToLower(choice) == "y" {
		fmt.Println("Enter the absolute filepath of the csv import file: ")
		fmt.Scanln(&Fp)
	}

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
