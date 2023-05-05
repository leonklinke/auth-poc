package api

import "net/http"

type api struct {
	router Router
}

func NewAPI(router Router) *api {
	return &api{router: router}
}

func (a *api) Start(addr string) error {
	var routes = RouterGroup{}
	routes = append(routes, homeRoutes...)
	a.addRoutes(routes)

	return a.router.Run(addr)
}

func (a *api) addRoutes(routes []Route) {
	for _, route := range routes {
		handler := route.HandlerFunc
		if len(route.PreHandlers) > 0 {
			for _, preHandler := range route.PreHandlers {
				handler = preHandler(handler)
			}
			handler = Sanitizer(handler)
		}
		a.router.AddRoute(route.Method, route.Path, parseHandler(handler))
	}
}

func parseHandler(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &Response{w, nil, true}
		handler(r, response)
	}
}
