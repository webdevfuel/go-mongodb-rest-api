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
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	ok := make(chan bool)
	addr := "localhost:3000"
	go http.ListenAndServe(addr, router)
	log.Printf("Server is listening on %s", addr)
	<-ok
}
