package main

import (
	"cli-app/app"
	"log"
	"os"
)

func main() {
	application := app.StartApp()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
