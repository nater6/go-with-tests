package main

import "fmt"

const MessagePrefix = "Hello, "
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return MessagePrefix + name
}

func main() {
	fmt.Println(Hello("Nate"))
}