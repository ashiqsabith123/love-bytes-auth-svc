package main

import (
	"log"

	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/ashiqsabith123/auth-svc/pkg/db"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error while loading config", err)
	}
	db.ConnectToDatabase(config)
}
