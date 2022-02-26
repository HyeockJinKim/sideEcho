package stats

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Stats interface {
	IncreaseRequestCount()
	IncreaseSuccessRequestCount()
	IncreaseFailureRequestCount()
	Clear()
}

type stats struct {
	requestCount        prometheus.Gauge
	successRequestCount prometheus.Gauge
	failureRequestCount prometheus.Gauge
}

func NewStats() Stats {
	return &stats{
		requestCount: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "request_count",
			Help: "request count of sideEcho",
		}),
		successRequestCount: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "success_request_count",
			Help: "success request count of sideEcho",
		}),
		failureRequestCount: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "failure_request_count",
			Help: "failure request count of sideEcho",
		}),
	}
}

func (s *stats) Clear() {
	s.requestCount.Set(0)
	s.successRequestCount.Set(0)
	s.failureRequestCount.Set(0)
}

func (s *stats) IncreaseRequestCount() {
	s.requestCount.Inc()
}

func (s *stats) IncreaseSuccessRequestCount() {
	s.successRequestCount.Inc()
}

func (s *stats) IncreaseFailureRequestCount() {
	s.failureRequestCount.Inc()
}
