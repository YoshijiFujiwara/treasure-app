package main

import (
	"github.com/joho/godotenv"
	server "github.com/YoshijiFujiwara/u22/backend"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file. %s", err)
	}

	dataSource := os.Getenv("DATABASE_DATASOURCE")
	dbDriver := os.Getenv("DB_DRIVER")
	if dataSource == "" {
		log.Fatal("Cannot get dataSource for database.")
	}

	s := server.NewServer()
	s.Init(dbDriver, dataSource)
	s.Run(os.Getenv("PORT"))
}
