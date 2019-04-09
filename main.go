package main

import (
	"flag"
	"fmt"
	"github.com/mscraftsman/devcon-feedback/store"
	"log"
	"net/http"
	"os"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/sessionize"
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
	_ = store.Init()
	router := initRouter()

	log.Println("Listening on :" + config.HTTPPort)
	_ = http.ListenAndServe(":"+config.HTTPPort, router)
}
