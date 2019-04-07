package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	//EnvironmentProd indicates Prod mode
	EnvironmentProd = "prod"

	//EnvironmentDev indicates Dev mode
	EnvironmentDev = "dev"
)

var (
	//Env is either prod (default) or dev
	Env string

	//BaseURL is the application base url including the protocol, excluding trailing slash
	BaseURL string

	//BoltDBPath is the path where to store the BoltDB database file. Must be writeable.
	BoltDBPath string

	//MeetupKey is used to authenticate via meetup.com
	MeetupKey string

	//MeetupSecret is used to authenticate via meetup.com
	MeetupSecret string

	//JWTSecret is used to encrypt the jwt token for cookies
	JWTSecret string

	//HTTPPort is the http port to listen to, defaults to 1337
	HTTPPort string

	//FrontURL is the url of the front web app
	FrontURL string
)

type loader struct {
	errors []error
}

func (l *loader) load(name string, def ...string) string {
	val := os.Getenv(name)
	if val != "" {
		return val
	} else if len(def) == 0 {
		l.errors = append(l.errors, fmt.Errorf("[env] environment variable not set: %s", name))
	} else {
		return def[0]
	}

	return ""
}

//Load sets up configuration for the application
func Load() {
	var l loader

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	Env = l.load("ENV", EnvironmentProd)
	BaseURL = l.load("BASE_URL")
	BoltDBPath = l.load("BOLT_DB_PATH")
	MeetupKey = l.load("MEETUP_KEY")
	MeetupSecret = l.load("MEETUP_SECRET")
	JWTSecret = l.load("JWT_SECRET")
	HTTPPort = l.load("HTTP_PORT", "1337")
	FrontURL = l.load("FRONT_URL")

	if len(l.errors) != 0 {
		for i := range l.errors {
			log.Println(l.errors[i])
		}
		os.Exit(1)
	}
}
