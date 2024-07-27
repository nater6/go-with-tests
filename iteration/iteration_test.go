package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("5 characters", func(t *testing.T) {
		expected := Repeat("a", 5)
		got := "aaaaa"
		assertCorrectMessage(t, got, expected)
	})
	t.Run("3 characters", func(t *testing.T) {
		expected := Repeat("y", 3)
		got := "yyy"
		assertCorrectMessage(t, got, expected)

	})
}

func BenchmarkRepeat(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("x", 5)
	fmt.Println(str)
	// Output: xxxxx
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
