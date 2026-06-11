# Go Basics

## Packages in Go

In Go, package is used to organize code into namespaces. Every Go file must belong to a package. Every folder falls under one main that is executable package. We can not create multiple.

It declares that this file is part of the main package. A main package is executable — it can be compiled into a standalone program. The main package must contain a func main() function, which is the entry point (where the program starts executing).

## The fmt Package

"fmt" is GO's formatting package from the standard library. It provides functions for formatted input and output (printing and scanning). (You can say I/O header files in c/c++ that we use)

### Common fmt functions:

- `fmt.Println()` — prints text with a newline at the end (what you're using)
- `fmt.Print()` — prints text without a newline
- `fmt.Printf()` — prints formatted text (like C's printf)
- `fmt.Scanf()` — reads user input

In our code, `fmt.Println("hello world")` takes the string and prints it to the console followed by a newline.

Without importing fmt, you'd have to use the built-in `println()` function instead, which is simpler but less flexible. The fmt package gives you more control over formatting.

## Running Go Programs

### Note:
1. In Go, to run the program we use "go run hello.go". Go tool compiles the source file into a temporary exe file which we can not see (happens in background) then immediately runs that temp binary and after program finishes temporary exe file is removed and we can see output right away. This is actually slower than "go build hello.go" because Go tool compiles the source file into real executable. It writes the binary disk in the current directory (hello.exe). It doesn't run the program. We can run the produced binary later with "./hello" or "./hello.exe" and that binary is platform independent and faster than "go run".

### Key difference
- `go run` = build + execute + cleanup
- `go build` = build only

## Go Background

Go is considered as one of the fastest-growing languages, especially for building scalable backend systems. It was created at Google beginning in 2007 as a simpler, more productive alternative to complex systems languages like C and C++. The designers wanted a language with fast compiled performance, strong support for concurrency, and much easier syntax so developers could build reliable server software more quickly.

In Go, all .go files inside the same folder belong to the same package. Go compiles all files in a folder together as one package. Multiple files can exist, but they share the same namespace. However, only one main() function is allowed in the main package, because it serves as the single entry point of the program. If more than one main() function is defined across files in the same folder, Go throws a main redeclared error.
