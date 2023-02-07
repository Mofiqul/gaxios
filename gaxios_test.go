package gaxios_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mofiqul/gaxios"
	"github.com/mofiqul/gaxios/utils"
)

type ResponseData struct {
	Message string `json:"message"`
}

var axios *gaxios.GAxios

func init() {
	axios = gaxios.New(
		&gaxios.GAxiosConfig{
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
	)
}

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
	t.Run("Test GET with instance", func(t *testing.T) {
		res, err := axios.Get(fmt.Sprintf("%s/test", server.URL))

		defer res.Data.Close()
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
	})

	t.Run("Test GET without instance", func(t *testing.T) {
		res, err := gaxios.Get(fmt.Sprintf("%s/test", server.URL),
			&gaxios.GAxiosConfig{
				Header: http.Header{
					"Accept": []string{"application/json"},
				},
			},
		)

		defer res.Data.Close()
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
	})
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

	t.Run("Test POST with instance", func(t *testing.T) {
		res, err := axios.Post(
			fmt.Sprintf("%s/test", server.URL),
			RequestBody{Name: "John Doe"},
		)

		defer res.Data.Close()
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
	})

	t.Run("Test POST without instance", func(t *testing.T) {
		res, err := gaxios.Post(
			fmt.Sprintf("%s/test", server.URL),
			RequestBody{Name: "John Doe"},
			&gaxios.GAxiosConfig{
				Header: http.Header{
					"Accept": []string{"application/json"},
				},
			},
		)

		defer res.Data.Close()
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
	})
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

	t.Run("Test PATCH with instance", func(t *testing.T) {
		res, err := axios.Patch(
			fmt.Sprintf("%s/test/2", server.URL),
			RequestBody{Name: "John Doe"},
		)

		defer res.Data.Close()
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
	})

	t.Run("Test PATCH without instance", func(t *testing.T) {
		res, err := gaxios.Patch(
			fmt.Sprintf("%s/test/2", server.URL),
			RequestBody{Name: "John Doe"},
			&gaxios.GAxiosConfig{
				Header: http.Header{
					"Accept": []string{"application/json"},
				},
			},
		)

		defer res.Data.Close()
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
	})
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

	t.Run("Testing DELETE with instance", func(t *testing.T) {
		res, err := axios.Delete(fmt.Sprintf("%s/test/2", server.URL))

		defer res.Data.Close()
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
	})

	t.Run("Testing DELETE without instance", func(t *testing.T) {
		res, err := gaxios.Delete(fmt.Sprintf("%s/test/2", server.URL),
			&gaxios.GAxiosConfig{
				Header: http.Header{
					"Accept": []string{"application/json"},
				},
			},
		)

		defer res.Data.Close()
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
	})
}

func TestCustomRoundTripper(t *testing.T) {
	mockRoundTripper := &utils.MockRoundTripper{
		RtFunc: func(r *http.Request) (*http.Response, error) {
			if r.URL.Path != "/roundtrip" {
				t.Errorf("Expected path: roundtrip, got: %s", r.URL.Path)
			}

			return &http.Response{StatusCode: 200}, nil
		},
	}

	axios := gaxios.New(&gaxios.GAxiosConfig{
		Transport: mockRoundTripper,
	})

	res, err := axios.Get("http://example.com/roundtrip")

	if err != nil {
		t.Errorf("Expected err: nil, got: %s", err.Error())
	}

	if res.Status != 200 {
		t.Errorf("Expected Status: 200, got %d", res.Status)
	}
}
