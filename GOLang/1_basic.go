package main

// First we declare main as package name it will tell the program what package is it.
// program will look for package main to run program first.

import "fmt"

// fmt is used for basic input / output
// import key to import package to go's file
// func main() is a function main. what code you write in this block it will be run first.
// sab kuch main mei he hota hai
func main() {

	// output
	fmt.Println("Hello World")

	// declaring variables
	var num int
	var name string

	// declaring constants
	// cannot be changed
	const Pi = 3.14

	fmt.Println("Input a number and your name")

	// taking input use &
	fmt.Scanln(&num)
	fmt.Scanln(&name)
	fmt.Println("Hey", name, num*num, "is the square")
}

// go run filename.go --> to run
// go build filename.go --> makes executable
// ./filename  --> to execute

// madhur3u@ZorinMadhur:~/Desktop/Virtual OS Desktop App /GO Lang$ 		go build op_ip.go
// madhur3u@ZorinMadhur:~/Desktop/Virtual OS Desktop App /GO Lang$ 		./op_ip
