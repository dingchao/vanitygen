package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

func instrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
) ServiceMiddleware {
	return func(next Vanitygen) Vanitygen {
		return instrmw{requestCount, requestLatency, countResult, next}
	}
}

type instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	Vanitygen
}

func (mw instrmw) Getvanityaddress(coin string, prefix string) (output string, key string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getvanityaddress", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output,key, err = mw.Vanitygen.Getvanityaddress(coin, prefix)
	return
}

