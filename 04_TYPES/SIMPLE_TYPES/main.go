package main

import (
	"fmt"
)

func main() {
	// Boolean Types
	var isTrue bool = true
	var isFalse bool = false

	fmt.Printf("isTrue: %t\n", isTrue)   // Output: isTrue: true
	fmt.Printf("isFalse: %t\n", isFalse) // Output: isFalse: false

	// Numeric Types
	var uint8Val uint8 = 255
	var int64Val int64 = -9223372036854775808
	var float32Val float32 = 3.14
	var complex64Val complex64 = 2 + 3i

	fmt.Printf("uint8Val: %d\n", uint8Val)     // Output: uint8Val: 255
	fmt.Printf("int64Val: %d\n", int64Val)     // Output: int64Val: -9223372036854775808
	fmt.Printf("float32Val: %.2f\n", float32Val) // Output: float32Val: 3.14
	fmt.Printf("complex64Val: %v\n", complex64Val) // Output: complex64Val: (2+3i)

	// String Types
	var message string = "Hello, my name is David!"

	fmt.Printf("message: %s\n", message)   
	fmt.Printf("Length of message: %d\n", len(message)) 
}
