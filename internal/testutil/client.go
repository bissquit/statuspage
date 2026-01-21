// Package testutil provides testing utilities for integration tests.
package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Client is an HTTP client for testing API endpoints.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// NewClient creates a new test client.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// LoginAs authenticates and stores the token.
func (c *Client) LoginAs(t *testing.T, email, password string) {
	t.Helper()

	resp, err := c.POST("/api/v1/auth/login", map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		t.Fatalf("login request failed: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("login failed: status=%d body=%s", resp.StatusCode, body)
	}

	var result struct {
		Data struct {
			Tokens struct {
				AccessToken string `json:"access_token"`
			} `json:"tokens"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("decode login response: %v", err)
	}

	c.Token = result.Data.Tokens.AccessToken
}

// LoginAsAdmin logs in as admin@example.com.
func (c *Client) LoginAsAdmin(t *testing.T) {
	t.Helper()
	c.LoginAs(t, "admin@example.com", "admin123")
}

// LoginAsOperator logs in as operator@example.com.
func (c *Client) LoginAsOperator(t *testing.T) {
	t.Helper()
	c.LoginAs(t, "operator@example.com", "admin123")
}

// LoginAsUser logs in as user@example.com.
func (c *Client) LoginAsUser(t *testing.T) {
	t.Helper()
	c.LoginAs(t, "user@example.com", "user123")
}

// ClearToken removes the stored token.
func (c *Client) ClearToken() {
	c.Token = ""
}

// GET performs a GET request.
func (c *Client) GET(path string) (*http.Response, error) {
	return c.do("GET", path, nil)
}

// POST performs a POST request with JSON body.
func (c *Client) POST(path string, body interface{}) (*http.Response, error) {
	return c.do("POST", path, body)
}

// PATCH performs a PATCH request with JSON body.
func (c *Client) PATCH(path string, body interface{}) (*http.Response, error) {
	return c.do("PATCH", path, body)
}

// DELETE performs a DELETE request.
func (c *Client) DELETE(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil)
}

func (c *Client) do(method, path string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	return c.HTTPClient.Do(req)
}

// DecodeJSON decodes response body into v.
func DecodeJSON(t *testing.T, resp *http.Response, v interface{}) {
	t.Helper()
	defer func() { _ = resp.Body.Close() }()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		t.Fatalf("decode response: %v", err)
	}
}

// ReadBody reads and returns response body as string.
func ReadBody(t *testing.T, resp *http.Response) string {
	t.Helper()
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	return string(body)
}
