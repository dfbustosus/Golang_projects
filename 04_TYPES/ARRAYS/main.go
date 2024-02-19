package main

// Be careful, ARRAYS are single type AND FIXED LENGHT!!!
import "fmt"

func main() {
	// Array Types
	// Array of bytes
	var byteArray [5]byte
	byteArray[0] = 'H'
	byteArray[1] = 'e'
	byteArray[2] = 'l'
	byteArray[3] = 'l'
	byteArray[4] = 'o'

	// Array of integers
	var intArray [3]int
	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3

	// Array of floats
	var floatArray [3]float32
	floatArray[0] = 1.12
	floatArray[1] = 2.24
	floatArray[2] = 3.34

	// Array of strings
	var stringArray [2]string
	stringArray[0] = "Hello"
	stringArray[1] = "World"

	// Print array values
	fmt.Println("Byte Array:", byteArray)   // Output: Byte Array: [72 101 108 108 111]
	fmt.Println("Int Array:", intArray)     // Output: Int Array: [1 2 3]
	fmt.Println("Float Array:", floatArray)     // Output: Int Array: [1 2 3]
	fmt.Println("String Array:", stringArray) // Output: String Array: [Hello World]
}