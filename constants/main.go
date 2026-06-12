package main

import "fmt"

func main() {

	// Single constant
	const companyName = "Google"

	// Numeric constant
	const ageLimit = 18

	// Floating point constant
	const pi = 3.14159

	fmt.Println("Company:", companyName)
	fmt.Println("Minimum Age:", ageLimit)
	fmt.Println("Value of PI:", pi)

	// Uncommenting the line below will cause an error
	// ageLimit = 21

}
//Output

// Company: Google
// Minimum Age: 18
// Value of PI: 3.14159
//By convention, camelCase()  naming convention is used for local constants, while PascalCase is used for exported constants.
