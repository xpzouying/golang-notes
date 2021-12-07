package zykit

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewPromHTTPHandler() http.Handler {

	return promhttp.Handler()
}

type CounterOpts = prometheus.CounterOpts

type Counter struct {
	core prometheus.Counter
}

func (c *Counter) Inc() {
	c.core.Inc()
}

func NewCounter(opts CounterOpts) *Counter {

	counter := prometheus.NewCounter(opts)

	prometheus.MustRegister(counter)
	return &Counter{counter}
}
