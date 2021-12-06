package zykit

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewPromHTTPHandler() http.Handler {

	return promhttp.Handler()
}
