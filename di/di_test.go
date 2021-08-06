package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jason")

	got := buffer.String()
	want := "Hello, Jason"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
