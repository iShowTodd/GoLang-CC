package main

import "fmt"

func main() {
	fmt.Println("Hello, world")

	// fmt.Printf("Hello, %s!\n") go vet will help you get this kind of errors

	fmt.Printf("Hello, %s!\n", "world")
}
