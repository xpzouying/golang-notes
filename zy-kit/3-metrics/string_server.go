package main

import (
	"log"
	"strings"
	"zykit"
)

type simpleStringServer struct{}

func (svr simpleStringServer) Upper(s string) string {
	result := strings.ToUpper(s)

	log.Printf("simple_string_server: got=%s result=%s", s, result)
	return result
}

func (svr simpleStringServer) Count(s string) int {
	return len(s)
}

type instrumentMiddleware struct {
	requestCounter zykit.Counter

	svc StringService
}

func (mw *instrumentMiddleware) Upper(s string) string {
	mw.requestCounter.Inc()

	return mw.svc.Upper(s)
}

func (mw *instrumentMiddleware) Count(s string) int {
	mw.requestCounter.Inc()

	return mw.svc.Count(s)
}
