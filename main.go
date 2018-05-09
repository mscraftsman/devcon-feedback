package main

import (
	"github.com/gorilla/mux"
	"github.com/mscraftsman/devcon-feedback/app"
)

func main() {
	config := app.Bootstrap()
	app.ServeHttp(config.HttpPort, func(router *mux.Router) error {
		//routes registration goes here
		return nil
	})
}
