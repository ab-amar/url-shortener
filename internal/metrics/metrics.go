package metrics

import "sync/atomic"

type Metrics struct {
	RequestsTotal         atomic.Int64
	ShortenRequestsTotal  atomic.Int64
	RedirectRequestsTotal atomic.Int64
	NotFoundTotal         atomic.Int64
}

func New() *Metrics {
	return &Metrics{}
}

func (m *Metrics) IncRequestsTotal() {
	m.RequestsTotal.Add(1)
}

func (m *Metrics) IncShortenRequestsTotal() {
	m.ShortenRequestsTotal.Add(1)
}

func (m *Metrics) IncRedirectRequestsTotal() {
	m.RedirectRequestsTotal.Add(1)
}

func (m *Metrics) IncNotFoundTotal() {
	m.NotFoundTotal.Add(1)
}

func (m *Metrics) GetRequestsTotal() int64 {
	return m.RequestsTotal.Load()
}

func (m *Metrics) GetShortenRequestsTotal() int64 {
	return m.ShortenRequestsTotal.Load()
}

func (m *Metrics) GetRedirectRequestsTotal() int64 {
	return m.RedirectRequestsTotal.Load()
}

func (m *Metrics) GetNotFoundTotal() int64 {
	return m.NotFoundTotal.Load()
}
