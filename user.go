package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CreateUserPayload struct {
	Name         string `json:"name" validate:"required"`
	EmailAddress string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

func createUser(payload CreateUserPayload) {
	log.Printf("creating user with name: %s, email: %s, password: %s", payload.Name, payload.EmailAddress, payload.Password)
}

func (h *handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		decoderError(err, w)
		return
	}

	v, err := getValidator()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = v.Struct(payload)
	if err != nil {
		validatorError(err, w)
		return
	}

	createUser(payload)
}
