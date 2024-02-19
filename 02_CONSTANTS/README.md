# Constants in Go

Go programming language provides different types of constants to define fixed values that cannot be changed during the execution of a program. Constants are expressions with a fixed value.

### Boolean Constants
Boolean constants represent the two truth values, `true` and `false`. They are used in conditional statements, loops, and logic operations.

Example:
```go
const (
    EnableDebugMode = true
    EnableCache     = false
)
```

### Rune Constants
Rune constants represent Unicode code points. They are used when working with characters, strings, and text processing.

Example:
```go
const (
    RuneA = 'A'
    RuneZ = 'Z'
)
```

### Integer Constants
Integer constants represent whole numbers. They are widely used in programming for counting, indexing, and representing numerical quantities.

Example:
```go
const (
    MaxConnections = 100
    DaysInWeek     = 7
)
```

### Floating-Point Constants
Floating-point constants represent real numbers with fractional parts. They are used in computations involving decimal numbers, such as scientific calculations or financial applications.

Example:
```go
const Pi = 3.14159
const Gravity = 9.81
```

### Complex Constants
Complex constants represent complex numbers with both real and imaginary parts. They are used in mathematical computations involving complex arithmetic.

Example:
```go
const (
    RealPart      = 2.5
    ImaginaryPart = 3.7
)
complexNum := complex(RealPart, ImaginaryPart)
```

### String Constants
String constants represent sequences of characters enclosed within double quotes. They are used to store textual data.

Example:
```go
const Greeting = "Hello, world!"
const DatabaseName = "my_database"
```

These constants play crucial roles in various programming scenarios and are fundamental to Go programming language.
