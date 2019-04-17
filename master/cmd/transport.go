package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func makeSortingEndpoint(svc SortingService, proxy *string, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(sortingRequest)
		input := make(map[int][]int)
		input[0] = []int{123, 543, 34, 78, 23, 45}
		input[1] = []int{167, 5456, 544, 28, 3, 465}
		input[2] = []int{53, 543, 346, 758, 253, 4665}
		input[3] = []int{653, 573, 347, 8, 283, 445}
		input[4] = []int{1, 53, 434, 758, 2223, 145}
		highestNumber := []int{}
		for i := 0; i < 5; i++ {
			svc = proxyingMiddleware(context.Background(), *proxy, logger)(svc)
			v, err := svc.Sorting(input[i])
			if err != nil {
				return sortingResponse{v, err.Error()}, nil
			}
			highestNumber = append(highestNumber, v)
		}
		sort.Ints(highestNumber)

		return sortingResponse{highestNumber[len(highestNumber)-1], ""}, nil
	}
}

func decodeSortingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sortingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSortingResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response sortingResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type sortingRequest struct {
	S string `json:"s"`
}

type sortingInputRequest struct {
	S []int `json:"s"`
}

type sortingResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}
