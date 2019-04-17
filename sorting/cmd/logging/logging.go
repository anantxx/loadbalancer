package logging

import (
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/loadbalancer/sorting/pkg/service"
)

func LoggingMiddleware(logger log.Logger) service.ServiceMiddleware {
	return func(next service.SortingService) service.SortingService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	service.SortingService
}

func (mw logmw) Sorting(s []int) (output int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "sorting",
			"input", "Array",
			"output", strconv.Itoa(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.SortingService.Sorting(s)
	return
}
