package main

import "fmt"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		return SpanishHelloPrefix + name
	case french:
		return FrenchHelloPrefix + name
	}
	return EnglishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Nate", "english"))
}
