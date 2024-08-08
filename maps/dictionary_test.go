package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Defined Key", func(t *testing.T) {
		key, value := "test", "this is just a test"
		dictionary := Dictionary{key: value}
		got, _ := dictionary.Search(key)
		want := value
		assertStrings(t, got, want)
	})

	t.Run("undefined word", func(t *testing.T) {
		key := "undefined"
		dictionary := Dictionary{}
		_, err := dictionary.Search(key)
		want := ErrNotFound
		assertErrors(t, err, want)
	})
}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		key, value := "test", "this value is for testing"
		dictionary := Dictionary{}
		err := dictionary.Add(key, value)
		assertErrors(t, err, nil)

		want := value
		got, err := dictionary.Search(key)

		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		assertStrings(t, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		key, value := "test", "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)
		assertErrors(t, err, ErrWordExists)
		got, _ := dictionary.Search(key)
		assertStrings(t, got, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("predefined word", func(t *testing.T) {

		word := "test"
		definition := "this is the original definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is the new definition"
		err := dictionary.Update(word, newDefinition)
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is the original definition"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word, definition := "test", "this is a test value"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertErrors(t, err, ErrNotFound)

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
