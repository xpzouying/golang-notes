package zykit

import (
	"context"
	"log"
	"time"
)

func WithTimeUsed(next Endpoint) Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		defer func(begin time.Time) {
			log.Printf("time_used: %s", time.Since(begin))
		}(time.Now())

		return next(ctx, request)
	}
}
