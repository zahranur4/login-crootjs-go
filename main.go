package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Halo Zahra! Backend kamu sudah LIVE di Render dan berhasil menerima data!",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	http.HandleFunc("/login", loginHandler)
	
	http.ListenAndServe(":"+port, nil)
}