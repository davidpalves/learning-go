package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello is exported to tests
func Hello(name string) string {
	if name != "" {
		return englishHelloPrefix + name
	}

	return englishHelloPrefix + "World"
}

func main() {
	fmt.Println(Hello("David"))
	fmt.Println(Hello(""))
}
