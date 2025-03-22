package curl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Curl is a simple http client
type Curl struct {
	client  *http.Client      // http client
	baseURL string            // base url
	headers map[string]string // headers
}

// NewCurl creates a new Curl instance
func NewCurl(baseURL string, timeout time.Duration) *Curl {
	return &Curl{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: strings.TrimSuffix(baseURL, "/"),
		headers: make(map[string]string),
	}
}

// SetHeader sets a header
func (c *Curl) SetHeader(key, value string) {
	c.headers[key] = value
}

// buildRequest builds a http request
func (c *Curl) buildRequest(ctx context.Context, method, urlPath string, queryParams map[string]string, body io.Reader) (*http.Request, error) {
	// Build the full url
	fullURL := c.baseURL + urlPath
	if queryParams != nil {
		query := url.Values{}
		for key, value := range queryParams {
			query.Add(key, value)
		}
		fullURL += "?" + query.Encode()
	}

	// Create a new http request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// doRequest sends a http request
func (c *Curl) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Get sends a GET request
func (c *Curl) Get(ctx context.Context, urlPath string, queryParams map[string]string) (*http.Response, error) {
	req, err := c.buildRequest(ctx, http.MethodGet, urlPath, queryParams, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

// Post sends a POST request
func (c *Curl) Post(ctx context.Context, urlPath string, body []byte) (*http.Response, error) {
	req, err := c.buildRequest(ctx, http.MethodPost, urlPath, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

// Put sends a PUT request
func (c *Curl) Put(ctx context.Context, urlPath string, body []byte) (*http.Response, error) {
	req, err := c.buildRequest(ctx, http.MethodPut, urlPath, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

// Delete sends a DELETE request
func (c *Curl) Delete(ctx context.Context, urlPath string) (*http.Response, error) {
	req, err := c.buildRequest(ctx, http.MethodDelete, urlPath, nil, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

// ReadResponseBody reads the response body
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("close response body failed: %v\n", err)
		}
	}(resp.Body)
	return io.ReadAll(resp.Body)
}
