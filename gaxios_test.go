package gaxios_test

import (
	"github.com/mofiqul/gaxios"
	"testing"

	"github.com/stretchr/testify/require"
)

var axios = gaxios.New()

func TestGetRequest(t *testing.T) {
	_, err := axios.Get("https://reqres.in/api/users", nil)
	require.NoError(t, err)
}

func TestPatchRequest(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}
	user := User{
		Name: "John Doe",
		Job:  "Senio SE",
	}
	_, err := axios.Patch("https://reqres.in/api/users/2", user, nil)
	require.NoError(t, err)
}

func TestDeleteRequest(t *testing.T) {
	_, err := axios.Delete("https://reqres.in/api/users/2", nil)
	require.NoError(t, err)
}

func TestPostRequest(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}
	user := User{
		Name: "John Doe",
		Job:  "Senio SE",
	}
	_, err := axios.Post("https://reqres.in/api/users", user, nil)
	require.NoError(t, err)
}
