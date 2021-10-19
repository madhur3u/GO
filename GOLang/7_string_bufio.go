package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// we are using bufio to read a whole line even with spaces
	// using scanln we cannot read a string if it have spaces
	// made a scanner and scan using it
	// it gives input in 1st and error
	scn := bufio.NewReader(os.Stdin)
	str, err := scn.ReadString('\n')

	if err == nil {
	}

	// length of string will be one more than what we input
	// it has one '\n' extra
	fmt.Println(len(str))

	// removing last character using slicing
	str = str[:len(str)-1]
	fmt.Println(str, len(str))

	// program to print substrings
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {

			fmt.Println(str[i : j+1])
		}
	}
}
