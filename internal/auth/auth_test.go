package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey my-secret-key")
		key, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if key != "my-secret-key" {
			t.Fatalf("expected key %q, got %q", "my-secret-key", key)
		}
	})

	t.Run("missing header", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected error for missing header")
		}
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})

	t.Run("malformed header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer token")
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected error for malformed header")
		}
	})
}
