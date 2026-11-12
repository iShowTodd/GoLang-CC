package main

import "fmt"

func main() {
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
}
