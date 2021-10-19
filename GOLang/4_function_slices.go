package main

import "fmt"

func main() {

	fmt.Println("Input a number")

	var N int
	var size int
	fmt.Scanln(&N)
	// taking input in N

	var odd = make([]int, 0)
	// now we have a slice named 'odd' giving it to a function and append all odd values till N

	appendODD(N, &odd)
	// calling a function
	// pass by reference to append in same slice

	size = Display(odd)
	// here we dont need append so pass by value, passing a copy

	fmt.Println("Size is ", size)

	// removing an element at 5th position in the slice (4th index)
	// we use append and slicing from 0-3 and 4-n and combine them
	odd = append(odd[:4], odd[5:]...)

	fmt.Println("The list after removing 5th element")
	size = Display(odd)
	fmt.Println("Size is ", size)

	// to get index of any number in our array
	GetIndex(odd, 18)

}

// this function append odd elements from 1 - N in a slice, pass by reference
func appendODD(n int, odd *[]int) {

	for i := 1; i <= n; i++ {

		if i%2 == 1 {
			*odd = append(*odd, i)
			// syntax to append at last position
		}
	}
}

// the int return after bracket tell us the return type
func Display(slice []int) int {

	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], "  ")
	}
	fmt.Println()
	return len(slice)
}

// this function print the index of k in slice and also returns it
func GetIndex(slice []int, k int) int {

	for i := 0; i < len(slice); i++ {

		if slice[i] == k {
			fmt.Println("Value", k, "found at position", i+1)
			return i
		}
	}
	fmt.Println("Value", k, "not found")
	return (-1)
}
