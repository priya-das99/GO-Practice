// In Go, package is used to organize code into namespaces. Every Go file must belong to a package.
// It declares that this file is part of the main package
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

// It provides functions for formatted input and output (printing and scanning).(You can say I/O header files in c/c++ that we use)


//Common fmt functions:

// fmt.Println() — prints text with a newline at the end (what you're using)
// fmt.Print() — prints text without a newline
// fmt.Printf() — prints formatted text (like C's printf)
// fmt.Scanf() — reads user input

// In our code, fmt.Println("hello world") takes the string and prints it to the console followed by a newline.

// Without importing fmt, you'd have to use the built-in println() function instead,
// which is simpler but less flexible. The fmt package gives you more control over formatting.

//Note:
// 1.In Go, To run the program we user "go run hello.go" go toll compiles the souce file into a temporary exe
//file which we can not see (happend in background ) then immidiately runs that temp binary and after progarm finsishes
//temporary exe file is removed and we can see output right away .This is actually slower than "go build hello.go" beacuse
//Go tool compiles the source file into real executatble.It writes the binary disk in the current directory(hello.exe)
//it doesnt run the program
//we can run the produced binary later with "./hello or ./hello.exe" and that binary is platform independent and faster than " go run"

//Key difference
// go run = build + execute + cleanup
// go build = build only
