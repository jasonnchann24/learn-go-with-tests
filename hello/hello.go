package hello

import (
	"fmt"

	"github.com/jasonnchann24/learn-go-with-tests/consts"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixSwitcher(language) + name
}

func prefixSwitcher(language string) (prefix string) {
	switch language {
	case consts.Spanish:
		prefix = consts.SpanishHello
	case consts.French:
		prefix = consts.FrenchHello
	default:
		prefix = consts.EnglishHello
	}

	prefix += ", "
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
