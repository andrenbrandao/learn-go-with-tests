package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishWorld = "World"
const spanishWorld = "Mundo"
const frenchWorld = "le Monde"
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	return greetingPrefix(language) + greetingSuffix(name, language)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func greetingSuffix(name string, language string) string {
	if name == "" {
		switch language {
		case spanish:
			return spanishWorld
		case french:
			return frenchWorld
		default:
			return englishWorld
		}
	}

	return name
}

func main() {
	fmt.Println(Hello("World", ""))
}
