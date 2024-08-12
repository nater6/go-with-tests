package main

import (
	"fmt"
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}
func ConfigurableRacer(aURL, bURL string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(aURL):
		return aURL, nil

	case <-ping(bURL):
		return bURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %q and %q", aURL, bURL)
	}

}
