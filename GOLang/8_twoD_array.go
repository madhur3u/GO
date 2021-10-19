package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var a [3][4]int

	a[0][0] = 5
	a[1][2] += 4

	fmt.Println(len(a))    // this give us size of rows
	fmt.Println(len(a[0])) // this give us size of columns
}
