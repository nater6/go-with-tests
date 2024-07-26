package main

import "testing"

// Write a test

// Make the compiler pass

// Run the test, see that it fails and check the error message is meaningful

// Write enough code to make the test pass

// Refactor

func TestHello(t *testing.T) {
	t.Run("Valid string passed to Hello()", func(t *testing.T) {
		got := Hello("Nate","")
		want := "Hello, Nate"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Empty string passed to Hello()", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
