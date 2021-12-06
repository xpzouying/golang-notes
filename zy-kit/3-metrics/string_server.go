package main

import (
	"log"
	"strings"
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
