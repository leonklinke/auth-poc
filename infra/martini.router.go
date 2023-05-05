package infra

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
)

type MartiniRouter struct {
	*martini.ClassicMartini
}

func NewMartiniRouter() *MartiniRouter {
	return &MartiniRouter{martini.Classic()}
}

func (r *MartiniRouter) AddRoute(method string, pattern string, handlerFunc http.HandlerFunc) {
	r.Get(pattern, handlerFunc)
}

func (r *MartiniRouter) Use(handlerFunc http.HandlerFunc) {
	r.Use(handlerFunc)
}

func (r *MartiniRouter) Run(addr string) error {
	fmt.Println("I'm using martini router")
	return http.ListenAndServe(addr, r)
}
