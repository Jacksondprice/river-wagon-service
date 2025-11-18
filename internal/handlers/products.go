package handlers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListProducts(w http.ResponseWriter, _ *http.Request) {
	items := []Product{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
