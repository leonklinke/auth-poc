package api

import (
	"net/http"
	"reflect"
	"strings"
)

func Sanitizer(next HandlerFunc) HandlerFunc {
	return func(request *http.Request, response *Response) {
		next(request, response)

		authorization := request.Header.Get("Authorization")

		if response.Succeeded {
			body := response.Body
			bodyType := body.Reflection.Type()
			responseBodyType := reflect.New(bodyType).Elem()
			for i := 0; i < bodyType.NumField(); i++ {
				isFieldAllowedToSend := false

				field := bodyType.Field(i)
				permissions := strings.Split(field.Tag.Get("permissions"), ",")

				for _, permission := range permissions {
					if authorization == permission || authorization == "admin" {
						isFieldAllowedToSend = true
					}
				}

				if isFieldAllowedToSend {
					fieldValue := reflect.ValueOf(body.Content).FieldByName(field.Name)
					responseBodyType.FieldByName(field.Name).Set(fieldValue)
				} else {
					responseBodyType.FieldByName(field.Name).SetZero()
				}

			}
			responseBody := responseBodyType.Interface()
			response.JSON(responseBody)
		}
	}
}
