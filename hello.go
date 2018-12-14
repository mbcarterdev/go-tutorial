package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//This is a public function because it starts with capital H
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

//This is a private function because it starts with lowercase h
//Can be private because it is only called by another function in this file.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}