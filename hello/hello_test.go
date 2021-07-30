package hello

import (
	"testing"

	"github.com/jasonnchann24/learn-go-with-tests/helpers"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jason", "English")
		want := "Hello, Jason"

		helpers.AssertEquals(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		helpers.AssertEquals(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elo", "Spanish")
		want := "Hola, Elo"

		helpers.AssertEquals(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Thom", "French")
		want := "Bonjour, Thom"

		helpers.AssertEquals(t, got, want)
	})
}
