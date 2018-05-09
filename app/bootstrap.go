package app

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
	// bootstrapped is a flag to prevent multiple bootstrapping
	bootstrapped = false

	// Env indicates in which environment (prod / dev) the application is running
	Env string

	// BaseURL Base URL of the webservice
	BaseURL string

	// MeetupRedirectURL Meetup.com Redirect URL
	MeetupRedirectURL string
)

// Config represents application configuration loaded during bootstrap
type Config struct {
	DB       *sql.DB
	HttpPort string

	// MeetupKey Meetup.com API Key
	MeetupKey string

	// MeetupSecret Meetup.com API Secret
	MeetupSecret string

	// JWT_SECRET JWT Secret Signing Token
	JWT_SECRET string
}

// Boostrap loads environment variables and initializes the application
func Bootstrap() *Config {
	var config Config

	if bootstrapped {
		return nil
	}

	godotenv.Load()

	Env = os.Getenv("ENV")
	if Env == "" {
		Env = EnvironmentProd
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("Environment variable DSN must be defined. Example: postgres://user:pass@host/db?sslmode=disable")
	}

	var err error
	config.DB, err = sql.Open("postgres", dsn)
	if err == nil {
		log.Println("Connected to database successfully.")
	} else if Env == EnvironmentDev {
		log.Println("Database connection failed: ", err)
	} else {
		log.Fatal("Database connection failed: ", err)
	}

	err = config.DB.Ping()
	if err == nil {
		log.Println("Pinged database successfully.")
	} else if Env == EnvironmentDev {
		log.Println("Database ping failed: ", err)
	} else {
		log.Fatal("Database ping failed: ", err)
	}

	config.HttpPort = os.Getenv("HTTP_PORT")
	if config.HttpPort == "" {
		config.HttpPort = "8888"
	}

	config.MeetupKey = os.Getenv("MEETUP_KEY")
	if config.MeetupKey == "" {
		log.Fatal("Environment variable MEETUP_KEY (Meetup.com API Key) must be defined.")
	}

	config.MeetupSecret = os.Getenv("MEETUP_SECRET")
	if config.MeetupSecret == "" {
		log.Fatal("Environment variable MEETUP_SECRET (Meetup.com API Secret) must be defined.")
	}

	config.JWT_SECRET = os.Getenv("JWT_SECRET")
	if config.JWT_SECRET == "" {
		log.Fatal("Environment variable JWT_SECRET (JWT Secret Signing Token) must be defined.")
	}

	BaseURL = os.Getenv("BASE_URL")
	if BaseURL == "" {
		log.Fatal("Environment variable BASE_URL (Base URL of the webservice) must be defined.")
	}

	MeetupRedirectURL = os.Getenv("MEETUP_REDIRECT_URL")
	if MeetupRedirectURL == "" {
		log.Fatal("Environment variable MEETUP_REDIRECT_URL (Meetup.com Redirect URL) must be defined.")
	}

	os.Clearenv() //prevent non-authorized access

	return &config
}
