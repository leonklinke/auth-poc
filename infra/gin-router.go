package infra

import (
	"fmt"
	"net/http"
	middleware "oauth-poc/infra/middlewares"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
}

func HighlineGinRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery(), middleware.Authentication(), middleware.Authorization())
	return engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{HighlineGinRouter()}
	// return &GinRouter{gin.Default()}

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
