package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jason")
	want := "Hello, Jason"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
