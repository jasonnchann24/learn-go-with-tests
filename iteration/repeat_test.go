package iteration

import (
	"fmt"
	"testing"

	"github.com/jasonnchann24/learn-go-with-tests/helpers"
)

func TestRepeat(t *testing.T) {
	got := Repeat("x", 5)
	want := "xxxxx"

	helpers.AssertEquals(t, got, want)
}

func ExampleRepeat() {
	repeated := Repeat("x", 3)
	fmt.Println(repeated)
	// Output: xxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 5)
	}
}
