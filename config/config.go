package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fluxynet/sequitur"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	//EnvironmentProd indicates Prod mode
	EnvironmentProd = "prod"

	//EnvironmentDev indicates Dev mode
	EnvironmentDev = "dev"

	//DatetimeLayout used by the system for session times
	DatetimeLayout = "2006-01-02T15:04:05"
)

var (
	//Tzone is the current time location
	Tzone *time.Location

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

	//AdditionalCORS contains a list of csv of cors urls that are allowed
	AdditionalCORS []string
)

var (
	//Now is the 'now' time to use when checking datetime for validation; used in dev environment
	Now *time.Time

	//VoteClosedAt is the timestring formatted as 2006-01-02T15:04:05
	VoteClosedAt *time.Time
)

var (
	errInvalidEnvVars = errors.New("invalid environment variables")
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

func (l *loader) loadDatetime(name string, def ...*time.Time) *time.Time {
	val := os.Getenv(name)

	if val == "" {
		if len(def) == 0 {
			l.errors = append(l.errors, fmt.Errorf("[env] environment variable not set: %s", name))
			return nil
		}
		return def[0]
	}

	if t, err := time.Parse(DatetimeLayout, val); err == nil {
		return &t
	}

	l.errors = append(l.errors, fmt.Errorf("[env] invalid datetime format (expecting %s) for variable: %s", DatetimeLayout, name))
	return nil
}

func (l *loader) loadLogLevel(def log.Level) log.Level {
	logLevel := l.load("LOG_LEVEL", "ERROR")
	switch logLevel {
	case "PANIC":
		return log.PanicLevel
	case "FATAL":
		return log.FatalLevel
	case "ERROR":
		return log.ErrorLevel
	case "WARN":
		return log.WarnLevel
	case "INFO":
		return log.InfoLevel
	case "DEBUG":
		return log.DebugLevel
	}

	if Env == EnvironmentProd {
		return log.ErrorLevel
	} else {
		return log.DebugLevel
	}
}

//Load sets up configuration for the application
func Load(filename string) {
	var (
		l        loader
		sequence sequitur.Sequence
	)

	if filename == "" {
		filename = ".env"
	}

	defer sequence.Catch(func(name string, err error) {
		if err == errInvalidEnvVars {
			for i := range l.errors {
				log.WithField("error", l.errors[i]).Error(name)
			}
		} else {
			log.WithField("error", err).Error(name)
		}
		os.Exit(1)
	})

	sequence.Do("loading env", func() error {
		return godotenv.Load(filename)
	})

	sequence.Do("reading environment variables", func() error {
		Env = l.load("ENV", EnvironmentProd)
		BaseURL = l.load("BASE_URL")
		BoltDBPath = l.load("BOLT_DB_PATH")
		MeetupKey = l.load("MEETUP_KEY")
		MeetupSecret = l.load("MEETUP_SECRET")
		JWTSecret = l.load("JWT_SECRET")
		HTTPPort = l.load("HTTP_PORT", "1337")
		FrontURL = l.load("FRONT_URL")
		SessionizeURL = l.load("SESSIONIZE_URL")
		Now = l.loadDatetime("NOW", nil)
		VoteClosedAt = l.loadDatetime("VOTE_CLOSED_AT")
		additionalCORS := l.load("ADDITIONAL_CORS", "_")

		if additionalCORS != "_" {
			AdditionalCORS = strings.Split(additionalCORS, ";")
		}

		if len(l.errors) == 0 {
			return nil
		}
		return errInvalidEnvVars
	})

	log.SetLevel(l.loadLogLevel(log.ErrorLevel))

	if t, err := time.LoadLocation("Indian/Mauritius"); err == nil {
		Tzone = t
	} else {
		log.Fatalln("Epic Failure when loading timezone")
	}
}
