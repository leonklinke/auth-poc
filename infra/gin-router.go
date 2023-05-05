package infra

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{gin.Default()}
}

func (r *GinRouter) AddRoute(method string, path string, handler http.HandlerFunc) {
	r.Engine.Handle(method, path, gin.WrapH(handler))
}

func (r *GinRouter) Use(handlerFunc http.HandlerFunc) {
	r.Engine.Use(gin.WrapH(handlerFunc))
}

func (r *GinRouter) Run(addr string) error {
	fmt.Println("I'm using gin router")
	return r.Engine.Run(addr)
}
