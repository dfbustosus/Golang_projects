# Understanding Types in Go

In Go, a type defines a set of values along with operations and methods specific to those values. Types play a fundamental role in Go programming, enabling developers to create structured and meaningful representations of data.

## Types in Go

A type in Go can be denoted by either a type name or a type literal. Let's break down the components of a type in Go:

- **Type Name**: A type may have a type name, which is typically an identifier representing the type. If the type is generic, the type name must be followed by type arguments enclosed in square brackets.

- **Type Literal**: Alternatively, a type can be specified using a type literal, which combines existing types to create a new type. This can include composite types such as arrays, structs, pointers, functions, interfaces, slices, maps, and channels.

## Named Types and Aliases

Go distinguishes between named types, defined types, and type parameters. Named types are predeclared types or types introduced with type declarations. An alias denotes a named type if the type given in the alias declaration is a named type. This distinction is important for understanding the behavior and usage of types in Go programs.

## Examples

Here are some examples of types in Go:

- **Predeclared Types**: These are types predeclared by the language, such as `int`, `string`, `bool`, etc.

- **Defined Types**: Types defined by users using type declarations. For example, `type MyInt int` creates a new type `MyInt` based on the `int` type.

- **Composite Types**: Types constructed using type literals, such as arrays, structs, pointers, functions, interfaces, slices, maps, and channels.

Understanding types is crucial for writing robust and maintainable Go code, as it enables developers to define clear data structures and enforce type safety in their programs.


# Arrays

Arrays in Go are a numbered sequence of elements of a single type, known as the element type. The number of elements in an array is fixed and is called the length of the array. Arrays are declared using the syntax `[length]elementType`, where `length` is a non-negative constant representable by a value of type `int` and `elementType` is the type of the elements in the array.

For example, to declare an array of integers with a length of 5:

```go
var numbers [5]int
```

# Simple Types

Simple types in Go refer to the basic built-in types provided by the language, such as `int`, `string`, `bool`, etc. These types are predeclared by the language and represent fundamental data types. Simple types are used to store individual values of specific kinds, such as integers, floating-point numbers, boolean values, characters, etc.

For example:
```go
var age int
var name string
var isStudent bool
```

# Slices

Slices in Go are dynamic, flexible data structures that provide a view into a contiguous segment of an underlying array. Unlike arrays, slices can grow and shrink dynamically during program execution. A slice is defined by specifying the type of elements it contains within square brackets `[]`.

```go
var numbers []int
```

Slices can be created using the make function or by slicing an existing array or another slice. Slices are highly versatile and widely used in Go for their ability to represent variable-length sequences of elements. They provide powerful capabilities for working with collections of data, such as appending, slicing, and modifying elements.

To increase the size of a slice, you can use the append function. Here's an example:

```go
floatSlice := make([]float64, 4)
floatSlice[0] = 3.245
floatSlice[1] = 1.23
floatSlice[2] = 3.212

// Increase the size of the slice by appending a new element
floatSlice = append(floatSlice, 4.567)
```

# Structs

A struct in Go is a way to group related data together. It's like a container that holds different pieces of information, each identified by a name. You can think of a struct as a blueprint for creating objects.


**Anatomy of a Struct**

- **Fields:** These are the individual pieces of data inside a struct. Each field has a name and a type. For example, a struct representing a person might have fields like name, age, and isMarried.

- **Declaration:** You define a struct using the type keyword followed by the name of the struct and a list of its fields enclosed in curly braces {}.

```go
type Person struct {
    name string
    age  int
    isMarried bool
}

// Creating an instance of Person struct
var john Person

// Setting values for fields
john.name = "John"
john.age = 30
john.isMarried = true

// Accessing values
fmt.Println(john.name)       // Output: John
fmt.Println(john.age)        // Output: 30
fmt.Println(john.isMarried)  // Output: true
```

A struct can contain another struct as one of its fields. This is called **embedding**.

```go
// Define Address struct
type Address struct {
    city  string
    state string
}

// Define Person struct with Address embedded
type Person struct {
    name    string
    age     int
    address Address  // Embedded struct
}

// Creating a person instance
var jane Person
jane.name = "Jane"
jane.age = 25
jane.address.city = "New York"
jane.address.state = "NY"
```

There are other technical aspects regaring structs but we will not go into them in for now.


# Pointers
Pointers in Go are variables that **store memory addresses**. They are widely used in Go for various purposes, including efficient memory management, passing references to functions, and working with data structures like slices and maps. Let's break down pointers in Go into simpler terms:

**Main properties**
- **Memory Address:** In Go, every variable has a memory address, which is a unique identifier for where it's stored in the computer's memory.
- **Pointer:** A pointer is a variable that stores the memory address of another variable.
- **Syntax:** To declare a pointer, you use the asterisk `*` symbol followed by the type of the variable it points to.

Here are some examples>
```go
var ptr *int
```

**Initializing Pointers**

- Zero Value: By default, when a pointer is declared, it is initialized to `nil`, which means it doesn't point to **any valid memory address**.

**Getting the Address of a Variable**

- Address-of Operator `(&)`: To get the memory address of a variable, you use the address-of operator `&`.

Here are examples:
```go
var num int = 42
var ptr *int = &num // ptr now points to the memory address of num
```

**Dereferencing Pointers**
- Dereference Operator `(*)`: To access the value stored at the memory address pointed to by a pointer, you use the dereference operator *.

Example:
```go
var num int = 42
var ptr *int = &num // referencing to num
fmt.Println(*ptr) // Output: 42 
```

**Passing Pointers to Functions**
- Pass by Reference: When you pass a pointer to a function, you're passing the memory address of the variable, allowing the function to modify the original value.

```go
func modifyValue(ptr *int) {
    *ptr = 100
}

var num int = 42
modifyValue(&num)
fmt.Println(num) // Output: 100
```