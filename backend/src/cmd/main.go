package main

import (
	"backend/src/internal/logs"
	"backend/src/internal/prometheus"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
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

	// Log an “info” with the path and timestamp
	log.WithFields(log.Fields{
		"path":     "/api/date",
		"datetime": now,
		"method":   r.Method,
		"remote":   r.RemoteAddr,
	}).Info("Served /api/date")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))

	// Log a health‐check event
	log.WithFields(log.Fields{
		"path":   "/health",
		"method": r.Method,
		"status": http.StatusOK,
	}).Info("Served /health")
}

func main() {
	prometheus.InitPrometheus() // Start Prometheus /metrics on :9091
	logs.InitLogs()             // Initialize JSON logging to /var/log/goapi/app.log

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/date", dateHandler)
	log.Info("Go API starting on :8080")
	http.ListenAndServe(":8080", nil)
}
