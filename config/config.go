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
)

type loader struct {
	errors []error
}

func (l loader) load(name string, variable *string) {
	*variable = os.Getenv(name)
	if *variable == "" {
		l.errors = append(l.errors, fmt.Errorf("[env] environment variable not set: %s", name))
	}
}

//Load sets up configuration for the application
func Load() {
	var l loader
	l.load("BOLT_DB_PATH", &BoltDBPath)
	l.load("MEETUP_KEY", &MeetupKey)
	l.load("MEETUP_SECRET", &MeetupSecret)
	l.load("JWT_SECRET", &JWTSecret)

	if len(l.errors) != 0 {
		for i := range l.errors {
			log.Println(l.errors[i])
		}
		os.Exit(1)
	}
}
