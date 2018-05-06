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
	EnvironmentDev = "DEV"
)

var (
	env      string
	db       *sql.DB
	httpPort string

	// MeetupKey Meetup.com API Key
	MeetupKey string

	// MeetupSecret Meetup.com API Secret
	MeetupSecret string
)

func bootstrap() {
	godotenv.Load()
	env = os.Getenv("ENV")

	if env == "" {
		env = EnvironmentProd
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("Environment variable DSN must be defined. Example: postgres://user:pass@host/db?sslmode=disable")
	}

	var err error
	db, err = sql.Open("postgres", dsn)
	if err == nil {
		log.Println("Connected to database successfully.")
	} else if env == EnvironmentDev {
		log.Println("Database connection failed: ", err)
	} else {
		log.Fatal("Database connection failed: ", err)
	}

	err = db.Ping()
	if err == nil {
		log.Println("Pinged database successfully.")
	} else if env == EnvironmentDev {
		log.Println("Database ping failed: ", err)
	} else {
		log.Fatal("Database ping failed: ", err)
	}

	httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8888"
	}

	// MeetupKey Meetup.com API Key
	MeetupKey = os.Getenv("MEETUP_KEY")
	if MeetupKey == "" {
		log.Fatal("Environment variable MEETUP_KEY (Meetup.com API Key) must be defined.")
	}

	// MeetupSecret Meetup.com API Secret
	MeetupSecret = os.Getenv("MEETUP_SECRET")
	if MeetupSecret == "" {
		log.Fatal("Environment variable MEETUP_SECRET (Meetup.com API Secret) must be defined.")
	}

}
