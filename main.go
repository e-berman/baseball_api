package main

import (
	"fmt"
	"log"
)

var Fp string

func Init() *DBPool {
	dbpool, err := NewDBPool()
	if err != nil {
		log.Fatal(err)
	}

	if err := dbpool.InitializePositionPlayerTable(); err != nil {
		log.Fatal(err)
	}
	if err := dbpool.InitializePitcherTable(); err != nil {
		log.Fatal(err)
	}

	return dbpool
}

func ImportPrompt(pool *DBPool, status bool) bool {
	var choice string

	fmt.Println("Do you want to import a csv file?: ")
	fmt.Println("[1]: Yes")
	fmt.Println("[2]: No")
	fmt.Scanln(&choice)

	if choice == "1" {
		fmt.Println("Are you importing pitcher data or position player data?: ")
		fmt.Println("[1]: Pitcher data")
		fmt.Println("[2]: Position Player data")
		fmt.Scanln(&choice)

		fmt.Println("Enter the absolute filepath of the csv import file: ")
		fmt.Scanln(&Fp)

		if choice == "1" {
			if err := pool.ImportPitcherDataFromCSV(); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := pool.ImportPositionPlayerDataFromCSV(); err != nil {
				log.Fatal(err)
			}
		}
	} else if choice == "2"{
		status = false
		return status
	} else {
		log.Println("Invalid Selection, please enter either 1 for Yes, or 2 for No")
	}
	return status
}

func main() {

	prompt_status := true

	dbpool := Init()
	defer dbpool.db.Close()
	
	for {
		prompt_status := ImportPrompt(dbpool, prompt_status)
		if prompt_status == false {
			break
		}
	}


	server := NewServer(":4242", dbpool)
	server.StartServer()
}
