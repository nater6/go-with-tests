package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("servers respond within deadline", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("got %v wanted no error", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("deadline expires", func(t *testing.T) {
		slowServer := makeDelayedServer(501 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(501 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, got := ConfigurableRacer(slowURL, fastURL, 500*time.Millisecond)
		if got == nil {
			t.Errorf("expected error but didn't get one")
		}
	})

}
