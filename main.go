package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/store"
)

// Version info
var (
	appVersion = "n/a"
	appCommit  = "n/a"
	appBuilt   = "n/a"
)

func main() {
	version := flag.Bool("v", false, "prints current app version")
	envfile := flag.String("c", ".env", "-c /path/to/my/env.file [defaults to .env]")

	flag.Parse()
	if *version {
		fmt.Printf("Version: %s \nCommit: %s \nBuilt: %s", appVersion, appCommit, appBuilt)
		os.Exit(0)
	}

	config.Load(*envfile)
	sessionize.LoadSessions()
	store.Init()
	router := initRouter()

	log.Println("Listening on :" + config.HTTPPort)
	http.ListenAndServe(":"+config.HTTPPort, router)
}

