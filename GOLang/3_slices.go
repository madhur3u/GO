package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var slice = make([]int, 0) // this is how we can make a slice
	// var arr[0]int this is how we make array
	// but we cannot change length of array
	// we cannot give a variable for [] also hence use SLICE

	for i := 0; i < 5; i++ {
		// to push a value to a slice
		slice = append(slice, i)
	}

	slice[2] += 5
	slice[3] += 5

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
