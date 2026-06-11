
// In Go, package is used to organize code into namespaces. Every Go file must belong to a package. every folder is  falls under one main that is execuatble package .we can not create multiple 
// It declares that this file is part of the main package
// A main package is executable — it can be compiled into a standalone program
// The main package must contain a func main() function, which is the entry point (where the program starts executing)

package main

import "fmt"
//"fmt " is GO's formatting package from the standard libary.
// It provides functions for formatted input and output (printing and scanning).(You can say I/O header files in c/c++ that we use)
func main() {
	fmt.Println("hello world")
}
