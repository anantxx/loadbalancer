package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/loadbalancer/sorting/pkg/service"
)

func MakeSortingEndpoint(svc service.SortingService, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(sortingRequest)
		s := req.S
		// s := []int{123, 543, 34, 78, 23, 45}
		v, err := svc.Sorting(s)
		if err != nil {
			return sortingResponse{v, err.Error()}, nil
		}
		return sortingResponse{v, ""}, nil
	}
}

func DecodeSortingRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
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
	S []int `json:"s"`
}

type sortingResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}
