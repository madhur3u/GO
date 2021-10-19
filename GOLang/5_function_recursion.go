package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)

	PrintDecreasingIncreasing(n)
}

// using recursion
// function is called inside the same function
func PrintDecreasingIncreasing(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	PrintDecreasingIncreasing(n - 1)
	fmt.Println(n)
}
