package main

import (
	"flag"
	"fmt"
	"gin-mvc/app/config"
	"gin-mvc/app/db"
	"gin-mvc/app/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	Env_load()

	config.Init(*environment)
	db.Init()
	server.Init()
}
