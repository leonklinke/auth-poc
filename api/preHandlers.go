package api

import (
	"net/http"
	"oauth-poc/model"
)

func AuthGuard(roles ...model.Role) PreHandler {
	return func(next HandlerFunc) HandlerFunc {
		return func(request *http.Request, response *Response) {
			authorization := request.Header.Get("Authorization")
			if authorization == "" {
				response.Succeeded = false
				response.WriteHeader(401)
				response.Write([]byte("Unauthorized"))
				return
			}
			next(request, response)
		}
	}
}
