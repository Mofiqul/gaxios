package gaxios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Header struct {
	Name  string
	Value string
}

type GAxiosConfig struct {
	Header  http.Header
	BaseUrl string
	Query   map[string]string
	Timeout time.Duration
}

type GAxios struct {
	client *http.Client
	config *GAxiosConfig
}

type GAxiosResponse struct {
	Status     int
	Data       io.Reader
	Header     http.Header
	StatusText string
	Request    *http.Request
}

func New(config *GAxiosConfig) *GAxios {
	gaxios := &GAxios{
		client: &http.Client{},
	}

	if config != nil {
		if config.Timeout != 0 {
			gaxios.client.Timeout = config.Timeout
		}
	}

	return gaxios
}

func processResponse(res *http.Response) (*GAxiosResponse, error) {
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Unable to read response body of bad response %w", err)
		}
		return nil, fmt.Errorf(
			"Response returned status with code %d: %+v, path: %s",
			res.StatusCode, string(body), res.Request.URL,
		)
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

func createRequest(
	method string, url string,
	payload interface{},
	cfg *GAxiosConfig,
) (*http.Request, error) {
	body, err := marshalBody(payload)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal payload %w", err)
	}

	if cfg != nil {
		if cfg.BaseUrl != "" {
			url = fmt.Sprintf("%s/%s", cfg.BaseUrl, url)
		}
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request %w", err)
	}

	if cfg != nil {
		if cfg.Header != nil {
			for k, v := range cfg.Header {
				req.Header.Set(http.CanonicalHeaderKey(k), strings.Join(v[:], ","))
			}
		}

		if cfg.Query != nil {
			for k, v := range cfg.Query {
				req.URL.Query().Add(k, v)
			}
		}
	}

	return req, nil
}

func (h *GAxios) Get(url string, cfg *GAxiosConfig) (resp *GAxiosResponse, err error) {
	req, err := createRequest(http.MethodGet, url, nil, cfg)
	if err != nil {
		return nil, err
	}
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}

func (h *GAxios) Patch(
	url string,
	payload interface{},
	cfg *GAxiosConfig,
) (*GAxiosResponse, error) {
	req, err := createRequest(http.MethodPatch, url, payload, cfg)
	if err != nil {
		return nil, err
	}
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}

func (h *GAxios) Post(
	url string,
	payload interface{},
	cfg *GAxiosConfig,
) (*GAxiosResponse, error) {
	req, err := createRequest(http.MethodPost, url, payload, cfg)
	if err != nil {
		return nil, err
	}
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}

func (h *GAxios) Delete(url string, cfg *GAxiosConfig) (*GAxiosResponse, error) {
	req, err := createRequest(http.MethodDelete, url, nil, cfg)
	if err != nil {
		return nil, err
	}
	res, err := h.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Unable perform request %w", err)
	}
	return processResponse(res)
}
