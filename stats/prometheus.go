package stats

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Stats interface {
	IncreaseRequestCount()
	Clear()
}

type stats struct {
	requestCount prometheus.Gauge
}

func NewStats() Stats {
	return &stats{
		requestCount: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "request_count",
			Help: "request count of sideEcho",
		}),
	}
}

func (s *stats) Clear() {
	s.requestCount.Set(0)
}

func (s *stats) IncreaseRequestCount() {
	s.requestCount.Inc()
}
