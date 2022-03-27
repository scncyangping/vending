package hello

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

func MakeUpperCaseEndpoints(service HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("MakeUpperCaseEndpoints")
		req := request.(uppercaseRequest)
		v, err := service.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, err
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(service HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		fmt.Println("MakeUpperCaseEndpoints")
		req := request.(countRequest)
		v := service.Count(req.S)
		return countResponse{v}, nil
	}
}

// transport中构造http.handler
func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// transport中构造http.handler
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
