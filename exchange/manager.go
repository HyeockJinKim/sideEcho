package exchange

import "sync/atomic"

type Manager interface {
	Buy(value uint64) error
	Sell(value uint64) error
}

type manager struct {
	value *uint64
}

func NewManager() Manager {
	return &manager{}
}

func (m *manager) Buy(value uint64) error {
	atomic.AddUint64(m.value, value)
	return nil
}

func (m *manager) Sell(value uint64) error {
	atomic.AddUint64(m.value, -value)
	return nil
}
