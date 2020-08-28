package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load("./auth.env")
	if err != nil {
		log.Fatal("Error loading auth.env file")
	}

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"), // environment variables
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_SERVER"),
		os.Getenv("APP_DB_NAME"))

	fmt.Println("Ingress backend has been initialized.")
	fmt.Println()
	a.Run(":8010")
}

func loadVars() {

}
