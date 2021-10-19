package main

import "fmt"

// this is how we can create a struct in go
// Student is the name of class/ struct
// we can access inner element by making a object
// s1 Student --> s1.name like this
type Student struct {
	name   string
	grades []int
	age    int
}

// function of a class when we use has to be used like this
// (s Student) this need to be aaded before function
// s here is object of Student and we can access it
func (s Student) Get_Percentage() float32 {

	p := 0

	// adding all grades
	// then divide by size of grades
	// and return
	for i := 0; i < len(s.grades); i++ {
		p = p + s.grades[i]
	}

	return (float32(p) / float32(len(s.grades)))
}

// here in this func we need to update age as provided in arguments
// for this we need to use pointer to update age in our main also
// if we do not use this, the age will get updated here but not be reflected in main
func (s *Student) Update_Age(x int) {
	s.age = x
}

// here we are updating our slice
// we are not using pointer here but still values get updated in main also
// pointer is not require in slice
func (s Student) Update_Grades(x int) {

	for i := 0; i < len(s.grades); i++ {

		if s.grades[i] < 90 {
			s.grades[i] += x
		}
	}
}

func main() {

	// here we made object s1 of type Student
	// and we are providing the data name, age etc to our struct
	s1 := Student{"Miku", []int{99, 89, 95, 87, 94, 90}, 25}
	s2 := Student{"Akari", []int{88, 89, 92, 86, 90, 96}, 25}

	// accessing from objects
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.grades[3])
	fmt.Println(s1.age)

	// how to use functions of a class
	fmt.Println(s1.Get_Percentage())

	// in this we need to update grades if they are less then 90, add argument
	s1.Update_Grades(5)
	s1.Update_Age(24)

	// print s1 again to see if values updated or not
	fmt.Println(s1)

	// can we use both objects together --> YES
	fmt.Println(s1.name, s2.name)
	fmt.Println(s2.Get_Percentage())
}
