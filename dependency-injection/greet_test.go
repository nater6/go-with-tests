package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "nater6"
	buf := bytes.Buffer{}
	Greet(&buf, name)
	want := "Hello, nater6"
	got := buf.String()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
