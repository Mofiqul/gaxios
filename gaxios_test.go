package gaxios_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mofiqul/gaxios"
)

type ResponseData struct {
	Message string `json:"message"`
}

var axios = gaxios.New(nil)

func TestGetMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			t.Errorf("Expected request to test, got: %s", r.URL.Path)
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "request success"}`))
	}))

	defer server.Close()
	res, err := axios.Get(fmt.Sprintf("%s/test", server.URL), &gaxios.GAxiosConfig{
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	})

	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err.Error())
	}

	resp := &ResponseData{}
	_ = json.NewDecoder(res.Data).Decode(resp)

	if res.Status != 200 {
		t.Errorf("Expected status code 200, got: %d", res.Status)
	}

	if resp.Message != "request success" {
		t.Errorf("Expected message to be request success, got: %v", resp.Message)
	}

}

func TestPostMethod(t *testing.T) {
	type RequestBody struct {
		Name string `json:"name"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			t.Errorf("Expected request to: test, got: %s", r.URL.Path)
		}

		if r.Method != http.MethodPost {
			t.Errorf("Expected request method: %s, got: %s", http.MethodPost, r.Method)
		}

		body := &RequestBody{}
		_ = json.NewDecoder(r.Body).Decode(body)

		if body.Name != "John Doe" {
			t.Errorf("Expected name in body param to be: John Doe, got: %s", body.Name)
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "request success"}`))
	}))

	defer server.Close()
	res, err := axios.Post(
		fmt.Sprintf("%s/test", server.URL),
		RequestBody{Name: "John Doe"},
		&gaxios.GAxiosConfig{
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
	)

	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err.Error())
	}

	resp := &ResponseData{}
	_ = json.NewDecoder(res.Data).Decode(resp)

	if res.Status != 201 {
		t.Errorf("Expected status code 200, got: %d", res.Status)
	}

	if resp.Message != "request success" {
		t.Errorf("Expected message to be request success, got: %v", resp.Message)
	}

}

func TestPatchMethod(t *testing.T) {
	type RequestBody struct {
		Name string `json:"name"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test/2" {
			t.Errorf("Expected request to: test/2, got: %s", r.URL.Path)
		}

		if r.Method != http.MethodPatch {
			t.Errorf("Expected request method: %s, got: %s", http.MethodPatch, r.Method)
		}

		body := &RequestBody{}
		_ = json.NewDecoder(r.Body).Decode(body)

		if body.Name != "John Doe" {
			t.Errorf("Expected name in body param to be: John Doe, got: %s", body.Name)
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "request success"}`))
	}))

	defer server.Close()
	res, err := axios.Patch(
		fmt.Sprintf("%s/test/2", server.URL),
		RequestBody{Name: "John Doe"},
		&gaxios.GAxiosConfig{
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
	)

	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err.Error())
	}

	resp := &ResponseData{}
	_ = json.NewDecoder(res.Data).Decode(resp)

	if res.Status != 200 {
		t.Errorf("Expected status code 200, got: %d", res.Status)
	}

	if resp.Message != "request success" {
		t.Errorf("Expected message to be request success, got: %v", resp.Message)
	}

}

func TestDeleteMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test/2" {
			t.Errorf("Expected request to test/2, got: %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("Expected request method: %s, got: %s", http.MethodDelete, r.Method)
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "request success"}`))
	}))

	defer server.Close()
	res, err := axios.Delete(fmt.Sprintf("%s/test/2", server.URL), &gaxios.GAxiosConfig{
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	})

	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err.Error())
	}

	resp := &ResponseData{}
	_ = json.NewDecoder(res.Data).Decode(resp)

	if res.Status != 200 {
		t.Errorf("Expected status code 200, got: %d", res.Status)
	}

	if resp.Message != "request success" {
		t.Errorf("Expected message to be request success, got: %v", resp.Message)
	}

}
