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

	// Declaring your slice

	/*
		Declaring a slice that might stay nil :
			var data []int
	*/

	/*
		You can create a slice using an empty slice literal:
		var x = []int{}
		This creates a slice with zero length and zero capacity. It is confusingly
		different from a nil slice. Because of implementation reasons,
		comparing a zero-length slice to nil returns false, while compar‐
		ing a nil slice to nil returns true. For simplicity, favor nil slices.
		A zero-length slice is useful only when converting a slice to JSON.
	*/

	x8 := []int{} // zero length and zero capacity
	fmt.Println(x8)

	/*
		Note : zero length != nil
			   nil slice == nil
	*/

	fmt.Println(x8 == nil) // false

	/*
		• If you are using a slice as a buffer,then specify a nonzero length.

		• If you are sure you know the exact size you want, you can specify the length and
		index into the slice to set the values. This is often done when transforming values
		in one slice and storing them in a second. The downside to this approach is that
		if you have the size wrong, you’ll end up with either zero values at the end of the
		slice or a panic from trying to access elements that don’t exist.

		• In other situations, use make with a zero length and a specified capacity. This
		allows you to use append to add items to the slice. If the number of items turns
		out to be smaller, you won’t have an extraneous zero value at the end. If the
		number of items is larger, your code will not panic.

	*/

	// Slicing to subslices
	list := []string{"a", "b", "c", "d"}

	slice1 := list[:2]
	slice2 := list[1:]
	slice3 := list[1:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// Never use Append with a subslice

	/*
		the subslice’s capacity
		is set to the capacity of the original slice, minus the starting offset of the subslice
		within the original slice. This means elements of the original slice beyond the end of
		the subslice, including unused capacity, are shared by both slices.
	*/

	slice4 := []string{"a", "b", "c", "d"}
	slice5 := slice4[:2]
	fmt.Println(cap(slice4), cap(slice5))
	slice5 = append(slice5, "z")
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)

	/*
		Even more confusing slices
			x := make([]string, 0, 5)
			x = append(x, "a", "b", "c", "d")
			y := x[:2]
			z := x[2:]
			fmt.Println(cap(x), cap(y), cap(z))
			y = append(y, "i", "j", "k")
			x = append(x, "x")
			z = append(z, "y")
			fmt.Println("x:", x)
			fmt.Println("y:", y)
			fmt.Println("z:", z)
	*/

	// The full slice expression protects against append
	slice6 := make([]string, 0, 5)
	slice6 = append(slice6, "a", "b", "c", "d")
	slice7 := slice6[:2:2]  // here you limited the capacity to be the length
	slice8 := slice6[2:4:4] //  here you limited the capacity to be the length
	// cap (high - low)
	fmt.Println(slice6)
	fmt.Println(slice7)
	fmt.Println(slice8)

	// Copy (independent copies)

	slice9 := []int{1, 2, 3, 4}
	slice10 := make([]int, 4)
	num := copy(slice10, slice9)
	fmt.Println(slice10, num)
	//  copy till you reach the length
	slice11 := []int{1, 2, 3, 4}
	slice12 := make([]int, 2)
	newSize := copy(slice12, slice11)
	fmt.Println(slice12, newSize)

	// copy from the middle of the source slice:
	slice13 := []int{1, 2, 3, 4}
	slice14 := make([]int, 2)
	copy(slice14, slice13[2:])

	// The copy function allows you to copy between two slices that cover overlapping
	// sections of an underlying slice
	slice15 := []int{1, 2, 3, 4}
	sameSize := copy(slice15[:3], slice15[1:])
	fmt.Println(slice15, sameSize)

	//	You can use copy with arrays by taking a slice of the array. You can make the array
	// either the source or the destination of the copy

	slice16 := []int{1, 2, 3, 4}
	arr5 := [4]int{5, 6, 7, 8}
	slice17 := make([]int, 2)
	copy(slice17, arr5[:])
	fmt.Println(slice17)
	copy(arr5[:], slice16)
	fmt.Println(arr5)
}
