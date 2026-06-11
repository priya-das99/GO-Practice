// Question 1: Basic Variable Declaration

// Write a Go program that:

// Creates a variable name and stores your name.
// Creates a variable age and stores your age.
// Prints both values.


//  Question 2: Swap Two Variables
// Task:

// Swap a and b without using a third variable.

package main

import "fmt"


func swapNumber() {
	a := 10
	b := 20

	a, b = b, a

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func main() {
	name := "Ankur"
	age := 20

	

	fmt.Println("My name is", name, "and age is", age)
	swapNumber()
}



