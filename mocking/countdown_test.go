package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := slySleeper{}

	Countdown(&buffer, &sleeper)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
