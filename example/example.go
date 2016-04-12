package main

import (
	"log"

	"github.com/dtylman/settings"
)

var conf = settings.New()

const confFile = "example.conf"

func main() {
	err := conf.Load(confFile, settings.FormatJSON)
	if err != nil {
		log.Printf("Cannot load configuration file: %v", err)
	}
	log.Printf("App bind: %s", conf.Get("app.bind", "no bind yet"))
	conf.Set("app.bind", "0.0.0.0:8080")
	err = conf.Save(confFile, settings.FormatJSON)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Saved to ", confFile)
	conf.Print(settings.FormatJSON)
}
