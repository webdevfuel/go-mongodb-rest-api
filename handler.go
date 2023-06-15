package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	client *mongo.Client
	db     string
}

func newHandler(client *mongo.Client, db string) handler {
	return handler{
		client: client,
		db:     db,
	}
}
