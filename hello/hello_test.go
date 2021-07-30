package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jason", "English")
		want := "Hello, Jason"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elo", "Spanish")
		want := "Hola, Elo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Thom", "French")
		want := "Bonjour, Thom"

		assertCorrectMessage(t, got, want)
	})
}
