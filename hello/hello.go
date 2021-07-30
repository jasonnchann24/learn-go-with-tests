package hello

import "fmt"

const (
	spanish      = "Spanish"
	french       = "French"
	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixSwitcher(language) + name
}

func prefixSwitcher(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}

	prefix += ", "
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
