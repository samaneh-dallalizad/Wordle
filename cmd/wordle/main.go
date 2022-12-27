package main

import (
	"wordleGame/internal/infrastructure"
	"wordleGame/internal/infrastructure/app"
)

func main() {

	//config setup
	infrastructure.ConfigSetup("config", ".")

	//setup

	srv := app.NewApplicationServer(nil)

	srv.State.HTTPServer.ListenAndServe()

}
