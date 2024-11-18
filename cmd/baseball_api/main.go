package main

import (
	"log"

	"github.com/e-berman/baseball_api/internal/routes"
	"github.com/e-berman/baseball_api/internal/db"
)

func main() {
	dbpool := setup()
	defer dbpool.Poolconn.Close()

	server := routes.NewServer(":4242", dbpool)
	server.StartServer()
}

func setup() *db.DBPool {
	dbpool, err := db.NewDBPool()
	if err != nil {
		log.Fatal(err)
	}

	if err := dbpool.InitializePositionPlayerTable(); err != nil {
		log.Fatal(err)
	}
	if err := dbpool.InitializePitcherTable(); err != nil {
		log.Fatal(err)
	}
	if err := dbpool.ImportPitcherDataFromCSV(); err != nil {
		log.Fatal(err)
	}
	if err := dbpool.ImportPositionPlayerDataFromCSV(); err != nil {
		log.Fatal(err)
	}

	return dbpool
}

