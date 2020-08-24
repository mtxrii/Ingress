package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, server, dbname string) {
	connectionString := "postgres://" + user + ":" + password + "@" + server + "/" + dbname + "?sslmode=disable"

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
