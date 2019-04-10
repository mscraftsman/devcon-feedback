package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/store"
	log "github.com/sirupsen/logrus"
)

// Version info
var (
	appVersion = "n/a"
	appCommit  = "n/a"
	appBuilt   = "n/a"
)

func main() {
	v := flag.Bool("v", false, "prints current app version")
	r := flag.Bool("r", false, "triggers database resetdb process")
	envfile := flag.String("c", ".env", "-c /path/to/my/env.file [defaults to .env]")

	flag.Parse()

	switch true {
	default:
		serve(*envfile)
	case *v:
		version()
	case *r:
		resetdb(*envfile)
	}
}

func serve(envfile string) {
	config.Load(envfile)
	sessionize.LoadSessions()

	if err := store.Init(); err != nil {
		log.WithField("error", err).Fatalln("serve:init:store")
	}

	router := initRouter()

	log.WithField("port", config.HTTPPort).Infoln("serve")
	_ = http.ListenAndServe(":"+config.HTTPPort, router)
}

func version() {
	fmt.Printf("Version: %s \nCommit: %s \nBuilt: %s", appVersion, appCommit, appBuilt)
	os.Exit(0)
}

func resetdb(envfile string) {
	var (
		response string
		sequence sequitur.Sequence
	)

	defer sequence.Catch(func(name string, err error) {
		log.WithField("error", err).Error("error when: " + name)
	})

	sequence.Do("asking for confirmation", func() error {
		fmt.Println("Are you sure you want to reset the database? [YES/...]")
		if _, err := fmt.Scanln(&response); err != nil {
			return err
		} else if response != "YES" {
			return errors.New("process cancelled")
		}
		return nil
	})

	sequence.Do("checking arithmetic skills", func() error {
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		fmt.Printf("%d + %d = ", x, y)

		if _, err := fmt.Scanln(&response); err != nil {
			return err
		} else if response != strconv.Itoa(x+y) {
			return errors.New("arithmetic skills bad")
		}

		return nil
	})

	sequence.Do("engaging countdown", func() error {
		fmt.Println("engaging countdown...")
		for x := 15; x > 0; x-- {
			fmt.Printf("...%d", x)
			time.Sleep(time.Second)
		}
		log.Println()
		log.Warn("Resetting database")

		return nil
	})

	sequence.Do("clearing db", func() error {
		config.Load(envfile)
		sessionize.LoadSessions()
		return store.ClearDB()
	})

	sequence.Then(func() {
		log.Println("feedback cleared from database")
	})
}
