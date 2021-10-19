package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)

	// this is how we store all return values when calling function
	a, b, c := Square_cube_message(n)
	println(a)
	println(b)
	println(c)

	d, e := anything(n)
	println(d)
	println(e)
}

// here we calculated a function
// this has 3 return values
// function in go can return multiplr values
// they can be of same or different types
func Square_cube_message(n int) (int, int, string) {

	sq := n * n
	cube := sq * n
	m := "Your square and cubes are calculated"

	// returning all values together
	return sq, cube, m

}

// Go's return values may be named.
// If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
func anything(n int) (x int, y int) {

	// defer inside a function is executed when it ends
	// that means it will 1st run the whole function
	// the going to run this
	defer fmt.Println("world <-- this was in defer")

	// defer can even be stacked, all defer will be stored in a stack
	// so last value to be stored, printed first
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Hello")

	x = n * 4 / 9
	y = n - x

	// provide a blank return as x, y will automatically return
	// A return statement without arguments returns the named return values
	// This is known as a "naked" return.
	return
}
