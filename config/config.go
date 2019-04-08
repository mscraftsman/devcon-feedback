package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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

	//SessionizeURL is used to get list of sessions
	SessionizeURL string

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
func Load(filename string) {
	var l loader

	if filename == "" {
		filename = ".env"
	}

	if err := godotenv.Load(filename); err != nil {
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
	SessionizeURL = l.load("SESSIONIZE_URL")
	logLevel := l.load("LOG_LEVEL", "ERROR")

	if len(l.errors) != 0 {
		for i := range l.errors {
			log.Println(l.errors[i])
		}
		os.Exit(1)
	}

	switch logLevel {
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	default:
		if Env == EnvironmentProd {
			log.SetLevel(log.ErrorLevel)
		} else {
			log.SetLevel(log.DebugLevel)
		}
	}
}
