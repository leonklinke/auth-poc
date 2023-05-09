package middleware

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthorizeClient(t *testing.T) {
	// authorizeSession
	req, err := http.NewRequest(http.MethodGet, "users/81de5fe8-eea1-11ed-a05b-0242ac120003", nil)
	require.NoError(t, err)

	headerAuth := "Bearer tokenClient"

	result, err := authorization(headerAuth, req)
	require.NoError(t, err)
	require.False(t, result) //client fails session validaiton
}

func TestAuthorizeMember(t *testing.T) {
	// authorizeSession
	req, err := http.NewRequest(http.MethodGet, "users/81de5fe8-eea1-11ed-a05b-0242ac120003", nil)
	require.NoError(t, err)

	headerAuth := "Bearer tokenMember"

	result, err := authorization(headerAuth, req)
	require.NoError(t, err)
	require.True(t, result) //Member success session validaiton
}
