package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type DateResponse struct {
	Date string `json:"date"`
}


func dateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	now := time.Now().Format("2006-01-02 15:04:05")
	json.NewEncoder(w).Encode(DateResponse{Date: now})
}


func main() {
	http.HandleFunc("/api/date", dateHandler)
	http.ListenAndServe(":8080", nil)
}


