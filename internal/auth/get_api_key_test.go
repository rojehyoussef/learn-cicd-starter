package auth

import (
	"net/http"
	"testing"
)

// GetAPIKey -
func TestGETAPIKey(t *testing.T) {
	headers := http.Header{}

	headers.Set("Authorization", "ApiKey abc123")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "abc123"
	if got != want {
		t.Fatalf("got: %q, want %q", got, want)
	}
}
