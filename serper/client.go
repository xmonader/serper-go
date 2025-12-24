package serper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const DefaultBaseURL = "https://google.serper.dev"

// ClientInterface defines the behavior of the Serper API client.
// This allows users to mock the client in their tests.
type ClientInterface interface {
	Search(ctx context.Context, req *Request) (*SearchResponse, error)
	Images(ctx context.Context, req *Request) (*ImageResponse, error)
	News(ctx context.Context, req *Request) (*NewsResponse, error)
	Videos(ctx context.Context, req *Request) (*VideoResponse, error)
	Places(ctx context.Context, req *Request) (*PlacesResponse, error)
}

// APIError represents an error returned by the Serper API.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("serper api error: %s (status %d)", e.Message, e.StatusCode)
}

// Client is the concrete implementation of ClientInterface.
type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// Option defines a functional option for configuring the Client.
type Option func(*Client)

// WithHTTPClient allows providing a custom http.Client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

// WithBaseURL allows providing a custom base URL (e.g., for a proxy).
func WithBaseURL(url string) Option {
	return func(c *Client) {
		if url != "" {
			c.baseURL = url
		}
	}
}

// WithTimeout is a convenience option to set a timeout on the default http.Client.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new Serper API client with the given API key and options.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: DefaultBaseURL,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Ensure Client implements ClientInterface.
var _ ClientInterface = (*Client)(nil)

func (c *Client) Search(ctx context.Context, req *Request) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.do(ctx, "/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Images(ctx context.Context, req *Request) (*ImageResponse, error) {
	var resp ImageResponse
	if err := c.do(ctx, "/images", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) News(ctx context.Context, req *Request) (*NewsResponse, error) {
	var resp NewsResponse
	if err := c.do(ctx, "/news", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Videos(ctx context.Context, req *Request) (*VideoResponse, error) {
	var resp VideoResponse
	if err := c.do(ctx, "/videos", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Places(ctx context.Context, req *Request) (*PlacesResponse, error) {
	var resp PlacesResponse
	if err := c.do(ctx, "/places", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) do(ctx context.Context, path string, payload interface{}, result interface{}) error {
	url := c.baseURL + path

	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var e struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&e); err == nil && e.Message != "" {
			return &APIError{StatusCode: resp.StatusCode, Message: e.Message}
		}
		return &APIError{StatusCode: resp.StatusCode, Message: "unknown error"}
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
