package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogs() {
	if err := os.MkdirAll("/var/log/goapi", 0o755); err != nil {
		log.Fatalf("could not create log directory: %v", err)
	}

	f, err := os.OpenFile("/var/log/goapi/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("could not open log file: %v", err)
	}

	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}
