package service

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var latencyExporter = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Namespace:  "be_k01",
		Subsystem:  "webapp",
		Name:       "latency",
		Help:       "recall latency in milliseconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"component", "status"},
)

var countExporter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "be_k01",
		Subsystem: "webapp",
		Name:      "count",
		Help:      "recall count",
	},
	[]string{"component", "type"},
)
