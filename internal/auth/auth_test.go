package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("valid header", func(t *testing.T) {
		header := http.Header{}
		want := "testkeytestkeytestkeytestkey"
		header.Set("Authorization", "ApiKey "+want)
		got, err := GetAPIKey(header)
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}

		if got != want {
			t.Errorf("invalid key. want %q, got %q", want, got)
		}
	})

	t.Run("invalid header", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "Bearer lksjdlkfjsdlkfj")
		_, err := GetAPIKey(header)
		if err == nil {
			t.Fatal("should have thrown an error")
		}

		if !errors.Is(err, ErrMalformedHeader) {
			t.Errorf("expected to throw malformed header error. got: %q", err.Error())
		}
	})

	t.Run("missing header", func(t *testing.T) {
		header := http.Header{}
		header.Set("Content-Type", "application/json")
		_, err := GetAPIKey(header)
		if err == nil {
			t.Fatal("should have thrown an error")
		}

		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Errorf("expected to throw no auth header error. got: %q", err.Error())
		}
	})
}
