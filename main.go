package main

import (
	"auth0-tutorial/app"
	"auth0-tutorial/pages"
	"auth0-tutorial/server"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
	err = app.Init()
	if err != nil {
		log.Panicln(err)
	}
	pages.Load("./pages/*.html")
}

func main() {
	router := mux.NewRouter()
	server.NewServer(router)
}
