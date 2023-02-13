package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const german = "German"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

// Upper case H to start the function name denotes that this is a public function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// lower case g to start the function name denotes that this is a private function
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// By providing a named return value in our function signature (prefix string)
	// we can simply use a blank "return"
	return
}

func main() {
	fmt.Println(Hello("world!", "English"))
}