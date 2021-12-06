package main

import (
	"log"
	"time"
)

type TimeUsedMiddleware struct {
	next StringService
}

func WithTimeUsedMiddleware(svc StringService) StringService {

	return &TimeUsedMiddleware{
		next: svc,
	}
}

func (mw *TimeUsedMiddleware) Upper(s string) (res string) {
	defer func(begin time.Time) {
		log.Printf("time_used: %s", time.Since(begin))

	}(time.Now())

	return mw.next.Upper(s)
}
