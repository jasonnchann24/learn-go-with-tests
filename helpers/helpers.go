package helpers

import (
	"testing"
)

func AssertEquals(t testing.TB, got, want interface{}) {
	switch want.(type) {
	case string:
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	case int:
		t.Helper()
		if got != want {
			t.Errorf("got '%d' but want '%d'", got, want)
		}
	default:
		t.Errorf("Assert type not available")
	}

}
