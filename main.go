package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("content-type", "application/json"))

	client := connectToDB()
	defer disconnectFromDB(client)

	db := "gomongo"
	h := newHandler(client, db)
	router.Get("/", h.helloWorld)

	ok := make(chan bool)
	addr := "localhost:3000"
	go http.ListenAndServe(addr, router)
	log.Printf("Server is listening on %s", addr)
	<-ok
}
