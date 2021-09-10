package presentation

import "io"

type HttpRequest struct {
	Body io.Reader
}

type HttpResponse struct {
	StatusCode int
	Body       interface{}
}

type Controller interface {
	Handle(req HttpRequest) HttpResponse
}
