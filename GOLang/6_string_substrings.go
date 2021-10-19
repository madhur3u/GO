package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	s := "abcd"

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {

			fmt.Println(s[i : j+1])
			// slicing string
		}
	}
}
