package main

import (
	"encoding/json"
	"meu_ip/middleware"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Content: r.RemoteAddr}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // optional but good practice
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloHandler)
	handler := middleware.EnableCORS(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}
