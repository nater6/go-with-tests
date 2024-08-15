package main

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("increment 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		want := 3
		assertCounter(t, counter, want)
	})

	t.Run("Safely use counter concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		for i := 0; i < wantedCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})

}
