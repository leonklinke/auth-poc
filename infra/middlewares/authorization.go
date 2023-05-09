package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	"github.com/google/uuid"
)

const BEARER_SCHEMA = "Bearer"

func authorization(header string, req *http.Request) (bool, error) {
	token := getToken(header)
	session, err := getSessionInfo(token)
	if err != nil {
		return false, err
	}

	if session.Role.IsAdmin() {
		return true, nil
	}

	sessionOk, err := authorize(session, req)
	if err != nil {
		return false, err
	}
	if !sessionOk {
		return false, nil
	}
	return true, nil
}

func getToken(header string) string {
	tokenString := header[len(BEARER_SCHEMA):]
	return tokenString
}

func authorize(session Session, req *http.Request) (bool, error) {
	path := req.URL.Path
	method := req.Method
	// /users/45s6df-4fs6d-5f4s/whatever...
	//    ^         ^
	//    |         |
	pices, err := getPathPices(path)
	if err != nil {
		return false, err
	}

	// users <- first part
	noun := pices[0]

	// 45s6df-4fs6d-5f4s <- second part
	subject, err := getPathSubject(pices)
	if err != nil {
		return false, err
	}

	rbcaOK, err := rbcaAuthorization(session, noun, method)
	if err != nil {
		return false, err
	}
	if !rbcaOK {
		return false, errors.New("role unauthorized")
	}

	return authorizationsMethods[noun](session, subject)
}

func getPathPices(path string) ([]string, error) {
	if len(path) == 0 {
		return nil, errors.New("empty path")
	}
	pices := strings.Split(path, "/")
	return pices, nil
}

func getPathSubject(pices []string) (string, error) {
	if len(pices) <= 1 {
		return "", errors.New("no route subject")
	}

	if !IsValidUUID(pices[1]) {
		return "", errors.New("invalid subject id")
	}

	return pices[1], nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func rbcaAuthorization(session Session, noun, method string) (bool, error) {
	enforcer, err := casbin.NewEnforcerSafe("./authorization_policy/model.conf", "./authorization_policy/policy.csv")
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	ok, err := enforcer.EnforceSafe(string(session.Role), noun, method)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("authorization error")
	}

	return ok, nil
}
