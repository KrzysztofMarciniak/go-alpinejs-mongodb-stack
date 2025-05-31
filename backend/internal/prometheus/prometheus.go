// prometheus.go
package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RequestCounter tracks the number of API requests
var RequestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"path"},
)

// InitPrometheus initializes Prometheus metrics and starts the /metrics endpoint on internal port 9091 only for prometheus server
func InitPrometheus() {
	prometheusMux := http.NewServeMux()
	prometheusMux.Handle("/metrics", promhttp.Handler())

	go func() {
		http.ListenAndServe("127.0.0.1:9091", prometheusMux)
	}()
}
