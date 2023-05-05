package infra

import (
	"fmt"
	"net/http"
)

type MuxRouter struct {
	*http.ServeMux
}

func NewMuxRouter() *MuxRouter {
	http.NewServeMux()
	return &MuxRouter{http.NewServeMux()}
}

func (m *MuxRouter) AddRoute(method string, path string, handler http.HandlerFunc) {
	m.HandleFunc(path, handler)
}

func (m *MuxRouter) Use(handlerFunc http.HandlerFunc) {
	//r.Engine.Use(gin.WrapH(handlerFunc))
}

func (m *MuxRouter) Run(addr string) error {
	fmt.Println("I'm using gin router")
	return http.ListenAndServe(addr, m)
}
