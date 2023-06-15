package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type person struct {
	Name string `bson:"name"`
}

func (h *handler) helloWorld(w http.ResponseWriter, _ *http.Request) {
	peopleCollection := h.client.Database(h.db).Collection("people")

	var personDocument person
	if err := peopleCollection.FindOne(context.Background(), bson.M{}).Decode(&personDocument); err != nil {
		w.Write([]byte(getGreeting(personDocument)))
		return
	}

	w.Write([]byte(getGreeting(personDocument)))
}

func getGreeting(p person) string {
	if p.Name != "" {
		return fmt.Sprintf("Hello, %s!", p.Name)
	}
	return "Hello, world!"
}
