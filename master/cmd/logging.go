package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next SortingService) SortingService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	SortingService
}

func (mw logmw) Sorting(s []int) (output int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "sorting",
			"input", "Array",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.SortingService.Sorting(s)
	return
}
