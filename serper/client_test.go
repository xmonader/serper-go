package serper

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Search(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("X-API-KEY") != "test-api-key" {
			t.Errorf("expected X-API-KEY header, got %s", r.Header.Get("X-API-KEY"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Check path
		if r.URL.Path != "/search" {
			t.Errorf("expected path /search, got %s", r.URL.Path)
		}

		// Mock response
		resp := SearchResponse{
			BaseResponse: BaseResponse{
				Credits: 1,
			},
			Organic: []OrganicResult{
				{Title: "Test Result", Link: "https://example.com"},
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	resp, err := client.Search(context.Background(), &Request{Q: "test"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Organic) != 1 {
		t.Errorf("expected 1 organic result, got %d", len(resp.Organic))
	}

	if resp.Organic[0].Title != "Test Result" {
		t.Errorf("expected title 'Test Result', got %s", resp.Organic[0].Title)
	}
}

func TestClient_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "invalid api key"}`))
	}))
	defer server.Close()

	client := NewClient("invalid-key", WithBaseURL(server.URL))
	_, err := client.Search(context.Background(), &Request{Q: "test"})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}

	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", apiErr.StatusCode)
	}

	if apiErr.Message != "invalid api key" {
		t.Errorf("expected message 'invalid api key', got %s", apiErr.Message)
	}
}

func TestOptions(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.example.com"
	
	client := NewClient("key", 
		WithHTTPClient(httpClient),
		WithBaseURL(baseURL),
	)

	if client.httpClient != httpClient {
		t.Error("WithHTTPClient option failed")
	}
	if client.baseURL != baseURL {
		t.Errorf("WithBaseURL option failed, got %s", client.baseURL)
	}
}
