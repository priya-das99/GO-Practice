# Constants in Go

A constant is a fixed value that cannot be changed after it is declared. Constants are declared using the `const` keyword and are generally used for values that remain the same throughout the program, such as PI, tax rates, or configuration values.

```go
const pi = 3.14159
```

## Types of Constants in Go

### 1. Typed Constants

A typed constant has an explicitly specified data type.

```go
const age int = 20
const name string = "Ankur"
const pi float64 = 3.14159
```

Note: The constant can only be used as the specified type.

### 2. Untyped Constants

An untyped constant does not have a fixed type until it is used in a context that requires one.

```go
const x = 10
const y = 3.14
```

Go automatically assigns an appropriate type when needed.

```go
var a int = x
var b float64 = y
```

Advantage: Untyped constants are more flexible and can be used with different compatible types.

## Naming Convention for Constants

Go follows the same naming rules for constants as variables:

### 1. camelCase (Most Common for Local Constants)

```go
const maxUsers = 100
const taxRate = 0.18
const pi = 3.14159
```

### 2. PascalCase (Exported Constants)

If a constant starts with a capital letter, it becomes **exported** (accessible from other packages).

```go
const Pi = 3.14159
const MaxUsers = 100
```

Example:

```go
// package mathutil
const Pi = 3.14159
```

```go
// another package
fmt.Println(mathutil.Pi)
```

### 3. ALL_CAPS (Not Common in Go)

Unlike C, C++, or Java, Go developers generally avoid this style.

```go
const MAX_USERS = 100 // Valid but not idiomatic Go
```

### Quick Rule

- `pi` → private/local to the package.
- `Pi` → exported/public to other packages.
- `PI` → valid but not preferred in Go style.

### Example

```go
package main

import "fmt"

const pi = 3.14159
const Pi = 3.14

func main() {
    fmt.Println(pi)
    fmt.Println(Pi)
}
```

### Note for Your Notes

> Constant names follow the same naming conventions as variables in Go. Go is case-sensitive, so `pi`, `Pi`, and `PI` are different identifiers. By convention, camelCase is used for local constants, while PascalCase is used for exported constants.

## Key Points

- Declared using the `const` keyword.
- Cannot be modified after declaration.
- Must be assigned a value at the time of declaration.
- Used for fixed values in a program.
- Can be Typed or Untyped.
- `iota` is a special identifier used to create a sequence of constants automatically.

```go
const (
	A = iota // 0
	B        // 1
	C        // 2
)