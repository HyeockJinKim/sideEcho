package stats

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestStats(t *testing.T) {
	promStats := NewStats()
	promStat, ok := promStats.(*stats)
	assert.True(t, ok)

	promStat.IncreaseRequestCount()
	assert.Equal(t, 1, int(testutil.ToFloat64(promStat.requestCount)))
	assert.Equal(t, 0, int(testutil.ToFloat64(promStat.successRequestCount)))
	assert.Equal(t, 0, int(testutil.ToFloat64(promStat.failureRequestCount)))
	t.Cleanup(func() {
		promStat.Clear()
	})
}

func TestStats_IncreaseSuccessRequestCount(t *testing.T) {
	promStats := NewStats()
	promStat, ok := promStats.(*stats)
	assert.True(t, ok)

	promStat.IncreaseSuccessRequestCount()
	assert.Equal(t, 0, int(testutil.ToFloat64(promStat.requestCount)))
	assert.Equal(t, 1, int(testutil.ToFloat64(promStat.successRequestCount)))
	assert.Equal(t, 0, int(testutil.ToFloat64(promStat.failureRequestCount)))
	t.Cleanup(func() {
		promStat.Clear()
	})
}

func TestStats_IncreaseFailureRequestCount(t *testing.T) {
	promStats := NewStats()
	promStat, ok := promStats.(*stats)
	assert.True(t, ok)

	promStat.IncreaseFailureRequestCount()
	assert.Equal(t, 0, int(testutil.ToFloat64(promStat.requestCount)))
	assert.Equal(t, 0, int(testutil.ToFloat64(promStat.successRequestCount)))
	assert.Equal(t, 1, int(testutil.ToFloat64(promStat.failureRequestCount)))
	t.Cleanup(func() {
		promStat.Clear()
	})
}
