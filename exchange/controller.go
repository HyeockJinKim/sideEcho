package exchange

import "sync/atomic"

//go:generate mockgen -package $GOPACKAGE -destination $PWD/exchange/mock_$GOFILE sideEcho/exchange Controller

type Controller interface {
	Buy(value uint64) error
	Sell(value uint64) error
}

type controller struct {
	value *uint64
}

func NewController() Controller {
	return &controller{
		value: new(uint64),
	}
}

func (m *controller) Buy(value uint64) error {
	atomic.AddUint64(m.value, value)
	return nil
}

func (m *controller) Sell(value uint64) error {
	atomic.AddUint64(m.value, -value)
	return nil
}
