package main

import "fmt"

func main() {
	// Arrays
	var arr [3]int // normal declaration with 3 default zero values
	fmt.Println(arr)

	arr2 := [3]int{10, 20, 30} // initialization with values

	fmt.Println(arr2)

	arr3 := [...]int{10, 20, 30} // automatically determine the size
	fmt.Println(arr3)

	arr4 := [12]int{1, 5: 4, 6, 10: 100, 15} //[1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15].
	fmt.Println(arr4)                        // index 5 takes a value 4 and index 10 takes the value of 100 , the others are zero

	size := len(arr3) // len determine the length of the array
	fmt.Println(size)

	// you can compare arrays with == and !=

	isEqual := arr2 == arr3
	fmt.Println(isEqual)

	// Slices (similar to arrays but flexible)
	var x []int
	fmt.Println(x == nil) // the only allowed comparison for slices

	x1 := []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x1) // similar to arrays

	var x2 [][]int  // simulation for 2-D slices
	fmt.Println(x2) // will produce nil (there are only few types that produce nil other than 0)

	// Slices Functions

	// append
	x3 := []int{1, 2, 3}
	x3 = append(x3, 4)
	fmt.Println(x3)
	y := []int{10, 20, 30, 40}
	x3 = append(x3, y...)
	fmt.Println(x3)

	// capacity
	/*
		the memory allocation rules of the Go Runtime
		1) Double it if it is less than 265
		2) if bigger (current + 768) / 4 and it converges at 25% growth
		3) 512 element has 63% growth
		4) 4096 has 30% growth
	*/

	var x4 []int
	fmt.Println(x4, len(x4), cap(x4))
	x4 = append(x4, 10)
	fmt.Println(x4, len(x4), cap(x4))
	x4 = append(x4, 20)
	fmt.Println(x4, len(x4), cap(x4))
	x4 = append(x4, 30)
	fmt.Println(x4, len(x4), cap(x4))
	x4 = append(x4, 40)
	fmt.Println(x4, len(x4), cap(x4))
	x4 = append(x4, 50)
	fmt.Println(x4, len(x4), cap(x4))

	// make → the only way of creating a slice with zero default value instead of nil
	// allows you to specify type, length and optionally the capacity

	x5 := make([]int, 5)
	fmt.Println(x5) // all are zeros

	x6 := make([]int, 5, 10) // you have now a capacity of 10
	fmt.Println(x6, len(x6), cap(x6))

	x7 := make([]int, 0, 10) // you can specify a 0 length slice with a determined capacity
	// Note you can not index any element
	fmt.Println(x7, len(x7), cap(x7))
	// Note you can not index any element
	x7 = append(x7, 4, 5, 6, 7)
	fmt.Println(x7, len(x7), cap(x7))

	// clear → it does only set the elements of slices to their zero values
	// the length is unchangeable using clear function
	clear(x7)
	fmt.Println(x7, len(x7))
}
