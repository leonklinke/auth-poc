package infra

import (
	"net/http"
)

type MockRouter struct {
	Routes      []string
	Middlewares []string
}

func (r *MockRouter) AddRoute(pattern string, handlerFunc http.HandlerFunc) {
	r.Routes = append(r.Routes, pattern)
}

func (r *MockRouter) Use(handlerFunc http.HandlerFunc) {
	r.Middlewares = append(r.Middlewares, "middleware")
}

func (r *MockRouter) Run(addr string) error {
	return nil
}

func NewMockRouter() *MockRouter {
	return &MockRouter{
		Routes:      []string{},
		Middlewares: []string{},
	}
}
