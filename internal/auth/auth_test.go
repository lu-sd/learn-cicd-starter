package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if key != "abc123" {
		t.Errorf("expected API key 'abc123', got '%s'", key)
	}
}

func TestGetAPIKey_NoHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer xyz")

	_, err := GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("expected 'malformed authorization header' error, got %v", err)
	}
}
