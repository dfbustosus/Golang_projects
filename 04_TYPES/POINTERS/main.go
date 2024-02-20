package main

import "fmt"

// Function to modify the value indirectly through a pointer
func modifyValue(ptr *int) {
    *ptr = 200
}

func main() {
    // Declare a variable num
    var num int = 42

    // Declare a pointer variable ptr
    var ptr *int

    // Assign the memory address of num to ptr
    ptr = &num

    // Print the memory address of num using ptr
    fmt.Println("Memory address of num:", ptr)

    // Dereference ptr to access the value stored at that memory address
    fmt.Println("Value of num (dereferencing):", *ptr)

    // Modify the value of num indirectly using ptr
    *ptr = 100

    // Print the new value of num
    fmt.Println("New value of num:", num)

    // Pass the pointer to a function to modify the original value
    modifyValue(ptr)

    // Print the updated value of num
    fmt.Println("Updated value of num:", num)
}
