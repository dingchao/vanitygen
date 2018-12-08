package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Vanitygen) Vanitygen {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	Vanitygen
}

func (mw logmw) Getvanityaddress(coin string, prefix string) (output string, key string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getvanityaddress",
			"input", coin,prefix,
			"output", output, key,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output,key, err = mw.Vanitygen.Getvanityaddress(coin, prefix)
	return
}

