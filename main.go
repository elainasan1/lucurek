package main

import (
	"fmt"
	"log"

	"filec2/attacks"
	"filec2/config"
	"filec2/db"
	"filec2/Handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := config.ReadConfig(); err != nil {
		log.Fatal("CFG", err)
	}

	fmt.Printf("Sky C2 Started | Version %s\n", config.Version)

	if err := db.Configure(); err != nil {
		log.Fatal("DB", err)
	}

	if err := db.SQL.Ping(); err != nil {
		log.Fatal("DB", err)
	}

	log.Printf(" [Connected To Database] [%s]", config.SQLHost)
	attacks.Configure()
	handlers.Serve()
}
