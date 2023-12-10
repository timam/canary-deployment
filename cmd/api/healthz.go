package main

import (
	"encoding/json"
	"net/http"
)

type Healthz struct {
	Status string `json:"status"`
}

func healthz(w http.ResponseWriter, r *http.Request) {
	response := Healthz{Status: "up"}
	json.NewEncoder(w).Encode(response)
}
