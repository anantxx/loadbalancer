package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/loadbalancer/sorting/pkg/service"
)

func instrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
) service.ServiceMiddleware {
	return func(next service.SortingService) service.SortingService {
		return instrmw{requestCount, requestLatency, countResult, next}
	}
}

type instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	service.SortingService
}

func (mw instrmw) Sorting(s []int) (output int, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sorting", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.SortingService.Sorting(s)
	return
}
