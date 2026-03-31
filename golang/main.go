package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Expose prometheus endpoint
func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}


