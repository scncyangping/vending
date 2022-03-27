package hello

import "strings"

type HelloService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type helloService struct{}

func (h *helloService) Uppercase(s string) (string, error) {
	return strings.ToUpper(s), nil
}

func (h *helloService) Count(s string) int {
	return len(s)
}

func NewHelloService() HelloService {
	return &helloService{}
}
