package main

import (
	"errors"
)

// SortingService provides operations on strings.
type SortingService interface {
	Sorting([]int) (int, error)
}

type sortingService struct{}

func (sortingService) Sorting(input []int) (int, error) {
	return 1, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for SortingService.
type ServiceMiddleware func(SortingService) SortingService
