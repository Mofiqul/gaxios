package gaxios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Header struct {
	Name  string
	Value string
}

type GAxiosConfig struct {
	Header http.Header
}

type GAxios struct {
	baseUrl string
	client  *http.Client
}

type GAxiosResponse struct {
	Status     int
	Data       io.Reader
	Header     http.Header
	StatusText string
	Request    *http.Request
}

func New() *GAxios {
	return &GAxios{
		client: &http.Client{},
	}
}

func processResponse(res *http.Response) (*GAxiosResponse, error) {
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Unable to read response body of bad response %w", err)
		}
		return nil, fmt.Errorf("Response returned status with code %d: %+v, path: %s", res.StatusCode, string(body), res.Request.URL)
	}
	resp := &GAxiosResponse{
		Status:     res.StatusCode,
		Header:     res.Header,
		Request:    res.Request,
		StatusText: res.Status,
		Data:       res.Body,
	}
	return resp, nil
}

func marshalBody(body interface{}) (io.Reader, error) {
	var payload io.Reader = nil
	if body != nil {
		b, err := json.Marshal(body)
		fmt.Print(string(b))
		if err != nil {
			return nil, err
		}
		payload = bytes.NewBuffer(b)
	}
	return payload, nil
}

func createRequest(method string, url string, payload interface{}) (*http.Request, error) {
	body, err := marshalBody(payload)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal payload %w", err)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request %w", err)
	}
	return req, nil
}

func (h *GAxios) Get(url string, config *GAxiosConfig) (resp *GAxiosResponse, err error) {
	req, err := createRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if config != nil {
		if config.Header != nil {
			for k, v := range config.Header {
				req.Header.Set(http.CanonicalHeaderKey(k), strings.Join(v[:], ","))
			}
		}
	}
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}

func (h *GAxios) Patch(url string, payload interface{}, config *GAxiosConfig) (*GAxiosResponse, error) {
	req, err := createRequest("PATCH", url, payload)
	if err != nil {
		return nil, err
	}
	if config != nil {
		if config.Header != nil {
			for k, v := range config.Header {
				req.Header.Set(http.CanonicalHeaderKey(k), strings.Join(v[:], ","))
			}
		}
	}
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}

func (h *GAxios) Post(url string, payload interface{}, config *GAxiosConfig) (*GAxiosResponse, error) {
	req, err := createRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	if config != nil {
		if config.Header != nil {
			for k, v := range config.Header {
				req.Header.Set(http.CanonicalHeaderKey(k), strings.Join(v[:], ","))
			}
		}
	}
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}

func (h *GAxios) Delete(url string, config *GAxiosConfig) (*GAxiosResponse, error) {
	req, err := createRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	if config != nil {
		if config.Header != nil {
			for k, v := range config.Header {
				req.Header.Set(http.CanonicalHeaderKey(k), strings.Join(v[:], ","))
			}
		}
	}
	res, err := h.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}
