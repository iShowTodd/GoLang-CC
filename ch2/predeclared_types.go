package main

import "fmt"

func main() {
	// Integer Literal
	fmt.Println(1_000_000)
	fmt.Println(0b10) // integer literal "binary"
	fmt.Println(0x10) // another integer literal "hexadecimal"

	// Floating Point
	fmt.Println(1.5) // floating point literal

	// Rune Literals
	fmt.Println('a')    // ASCII code for a
	fmt.Println('\141') // Octal representation of a
	fmt.Println('\x61') // Hexadecimal representation of a

	// String Literal
	fmt.Println("This is a String Literal")
}
