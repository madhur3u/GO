package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var num int

	fmt.Println("Input a number and we print all odd numbers till that")
	fmt.Scanln(&num)

	// for loop in GO
	for i := 1; i <= num; i++ {

		if i%2 == 1 {
			fmt.Print(i, "  ")
		}

	}
	fmt.Println()

	fmt.Println("Now even numbers :")

	// this is how we can write for in parts
	// just like while loop
	i := 2
	for i <= num {

		if i%2 == 0 {
			fmt.Print(i, "  ")
		}

		i++
	}
	fmt.Println()
}

// continue and break can be used just like other languages
