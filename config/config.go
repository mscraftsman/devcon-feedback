package config

import (
	"fmt"
	"log"
	"os"
)

var (
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
)

type loader struct {
	errors []error
}

func (l loader) load(name string, variable *string, def ...string) {
	*variable = os.Getenv(name)
	if *variable == "" {
		if len(def) == 0 {
			l.errors = append(l.errors, fmt.Errorf("[env] environment variable not set: %s", name))
		} else {
			*variable = def[0]
		}
	}
}

//Load sets up configuration for the application
func Load() {
	var l loader
	l.load("BASE_URL", &BaseURL)
	l.load("BOLT_DB_PATH", &BoltDBPath)
	l.load("MEETUP_KEY", &MeetupKey)
	l.load("MEETUP_SECRET", &MeetupSecret)
	l.load("JWT_SECRET", &JWTSecret)
	l.load("HTTP_PORT", &HTTPPort, "1337")

	if len(l.errors) != 0 {
		for i := range l.errors {
			log.Println(l.errors[i])
		}
		os.Exit(1)
	}
}
