package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("./auth.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"), // environment variables
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_SERVER"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}

func loadVars() {

}
