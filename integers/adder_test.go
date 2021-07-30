package integers

import (
	"fmt"
	"testing"

	"github.com/jasonnchann24/learn-go-with-tests/helpers"
)

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	want := 5

	helpers.AssertEquals(t, got, want)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
