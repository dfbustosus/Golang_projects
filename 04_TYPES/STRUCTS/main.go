package main

import (
	"fmt"
)

type Person struct {
    name string
    age  int
    isMarried bool
}

// Creating an instance of Person struct
var john Person

// Method associated with Person struct
func (p Person) sayHello() {
    fmt.Println("Hello, my name is", p.name)
}
//############################################
// Define Address struct (Embebbed Struct)
type Address struct {
    city  string
    state string
}

// Define Person struct with Address embedded
type Person2 struct {
    name    string
    age     int
    address Address  // Embedded struct
}

// Creating a person instance
var jane Person2

func main() {

// Setting values for fields (struct 1)
john.name = "John"
john.age = 30
john.isMarried = true

// Setting values for fields (struct 2)
jane.name = "Jane"
jane.age = 25
jane.address.city = "New York"
jane.address.state = "NY"


// First Struct
fmt.Println(john.name)       // Output: John
fmt.Println(john.age)        // Output: 30
fmt.Println(john.isMarried)  // Output: true
// Calling the method
john.sayHello() // Output: Hello, my name is John

// Second Struct
fmt.Println(jane.name)       // Output: John
fmt.Println(jane.age)        // Output: 30
fmt.Println(jane.address)  // Output: true


}
