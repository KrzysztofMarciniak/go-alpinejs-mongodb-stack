package main

import (
	"backend/internal/prometheus"
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

	// Prometheus: increment metric
	prometheus.RequestCounter.WithLabelValues("/api/date").Inc()

	now := time.Now().Format("2006-01-02 15:04:05")
	json.NewEncoder(w).Encode(DateResponse{Date: now})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	prometheus.InitPrometheus() // Start Prometheus server

	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/date", dateHandler)
	http.ListenAndServe(":8080", nil)
}
