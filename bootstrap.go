package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	//EnvironmentProd represents production environment
	EnvironmentProd = "PROD"

	//EnvironmentDev represents development environment
	EnvironmentDev  = "DEV"
)

var (
	env string
    db  *sql.DB
	
	//  
	 
	
)

func bootstrap() {
	var err error

	godotenv.Load()

	dsn := os.Getenv("DSN")
	env = os.Getenv("ENV")

	if env == "" {
		env = EnvironmentProd
	}

	if dsn == "" {
		log.Fatal("Environment variable DSN must be defined. Example: postgres://user:pass@host/db?sslmode=disable")
	}

	db, err = sql.Open("postgres", dsn)
	if err == nil {
		log.Println("Connected to database successfully.")
	} else if (env == EnvironmentDev) {
		log.Println("Database connection failed: ", err)
	} else {
		log.Fatal("Database connection failed: ", err)
	}

	err = db.Ping()
	if err == nil {
		log.Println("Pinged database successfully.")
	} else if (env == EnvironmentDev) {
		log.Println("Database ping failed: ", err)
	} else {
		log.Fatal("Database ping failed: ", err)
	}
}