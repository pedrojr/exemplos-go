package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
}

func main() {
	m := initializeMetrics()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.RequestCount.Add(1)
		time.Sleep(100 * time.Millisecond)
		duration := time.Since(start).Seconds()
		m.RequestLatency.Observe(duration)
		fmt.Fprintf(w, "pong")
	})

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server started on 8090")
	http.ListenAndServe(":8090", nil)
}

func initializeMetrics() *Metrics {
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "example_counter",
		Name:      "http_requests_total",
		Help:      "Total number of HTTP requests.",
	}, []string{})

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "example_summary",
		Name:      "http_request_duration_seconds",
		Help:      "Duration of HTTP requests in seconds.",
	}, []string{})

	return &Metrics{
		RequestCount:   requestCount,
		RequestLatency: requestLatency,
	}
}
