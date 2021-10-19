package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	// declaring a hashmap or dictionary, here simply called MAP
	// in 2nd line we are initializing our hashmap
	var mymap map[string]int
	mymap = make(map[string]int)

	// this is how we add key-value pais in MAP
	mymap["Miku"] = 24
	mymap["Rie"] = 27
	mymap["Akari"] = 26

	var key string
	var value int
	var N int

	// taking input from user
	fmt.Println("How many values to add ?")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {

		fmt.Println("Enter a VA Name (key) and age (value)")
		fmt.Scan(&key)
		fmt.Scan(&value)
		mymap[key] = value
	}

	fmt.Println()
	fmt.Println("Name", "\t\t", "Age")
	fmt.Println()

	// printing a MAP
	// key will store key of map
	// val will store value of the key
	// loop runs till it reads whole MAP
	for key, val := range mymap {
		fmt.Println(key, "\t\t", val)
	}
}
