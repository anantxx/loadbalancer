package service

import (
	"errors"
	"sort"
)

// SortingService provides operations on strings.
type SortingService interface {
	Sorting([]int) (int, error)
}

type SService struct{}

func (s2 SService) Sorting(nameList []int) (int, error) {
	sort.Ints(nameList)
	return nameList[len(nameList)-1], nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for SortingService.
type ServiceMiddleware func(SortingService) SortingService
