package main

import (
	"fmt"
	"math/cmplx"
)

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

	// Booleans

	var flag bool     // The Default value is zero
	isAwesome := true // var isAwesome = true "Due to I use lints and its working with best practices"

	fmt.Println(flag)
	fmt.Println(isAwesome)

	// Numeric Types

	var num1 int8 = 1
	var num2 int16 = 2
	var num3 int32 = 3
	var num4 int64 = 4
	var num5 uint8 = 5
	var num6 uint16 = 6
	var num7 uint32 = 7
	var num8 uint64 = 8
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
	fmt.Println(num8)

	var num9 float32 = 9.9
	fmt.Println(num9)

	// var num10 float64 = 10.999999999999999
	// fmt.Println(num10)

	// Complex  Types → Has 2 parts real + imaginary
	// (x1 + x2) (y1+y2) can be addition or any logical operation
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	// a glimpse of Strings  and Runes
	var myFirstInitial rune = 'A'   // rune is an alias for int32
	var mySecondInitial int32 = 'B' // bad practice use Rune instead

	// Explicit Type Conversion

	fmt.Println(myFirstInitial + mySecondInitial)
	var digit1 int = 10
	var digit2 float64 = 30.2
	var sum1 float64 = float64(digit1) + digit2
	var sum2 int = digit1 + int(digit2)
	fmt.Println(sum1, sum2)

	var digit3 int = 10
	var digit4 byte = 100
	var sum3 int = digit3 + int(digit4)
	var sum4 byte = byte(digit3) + digit4
	fmt.Println(sum3, sum4)

	// literals are untyped

	var ans float64 = 5 * 20.5
	fmt.Println(ans)

	// var vs :=

	var baseSalary int = 4000
	bonus := 500
	var overtime int

	var hoursWorked, daysOff int = 160, 2
	var rating, level int
	name, years := "Ahmed", 3

	var (
		department          string
		raise                       = 200
		taxRate             float32 = 0.15
		firstName, lastName         = "Ahmed", "Arafa"
		isActive, isManager bool
	)

	fmt.Println("Total pay:", baseSalary+bonus+overtime+raise)
	fmt.Println("Hours:", hoursWorked, "Days off:", daysOff)
	fmt.Println("Rating:", rating, "Level:", level)
	fmt.Println(name, years)
	fmt.Println(firstName, lastName)
	fmt.Printf("%q %.2f\n", department, taxRate)
	fmt.Println(isActive, isManager)

	name2 := "Ahmed"
	baseSalary2 := 4000
	bonus2 := 500.5
	isActive2 := true

	hoursWorked2, daysOff2 := 160, 2
	role2, level2 := "Backend", 3

	fmt.Println(name2, baseSalary2, bonus2, isActive2)
	fmt.Println(hoursWorked2, daysOff2, role2, level2)
	fmt.Printf("%T %T %T %T\n", name2, baseSalary2, bonus2, isActive2)

	baseSalary2 = baseSalary2 + 300
	fmt.Println(baseSalary2)

	daysOff2, vacation2 := 5, 10
	fmt.Println(daysOff2, vacation2)

	const immutableNum int64 = 10
	println(immutableNum)

	const (
	// I can declare here exactly like var
	)

	// immutableNum = 200 // will not compile

	// Go doesn’t provide a way to specify that a value calculated at runtime is immutable.

	// x := 5
	// y := 10
	// const z = x + y // this won't compile!

	// Untyped Const → acts exactly like a literal
	// literal = directly written value
	// untyped literal = a value that can be of more than 1 type
	const unTypedConstant = 10

	var typed1 int = unTypedConstant
	var typed2 float64 = unTypedConstant
	var typed3 byte = unTypedConstant

	fmt.Println(typed1)
	fmt.Println(typed2)
	fmt.Println(typed3)
}
