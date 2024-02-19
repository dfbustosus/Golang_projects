package main

import "fmt"
// Be careful, SLICES are single type BUT NOT FIXED LENGHT!!!
func main() {
	// Slice Types
	// Slice of integers
	intSlice := []int{1, 2, 3, 4, 5} // 1 way with [] type{elements}

	// Slice of strings
	stringSlice := make([]string, 3) // 2 way make([]type,#elements)
	stringSlice[0] = "apple"
	stringSlice[1] = "banana"
	stringSlice[2] = "orange"

	// Slice of floats
	floatSlice := make([]float64, 4) // 2 way make([]type,#elements)
	floatSlice[0] = 3.245;floatSlice[1] = 1.23;floatSlice[2] = 3.212 // Default the last one =0
	// Increase the  slice size
	floatSlice = append(floatSlice, 4.567) // It should be added at the end position

	// Print slice values
	fmt.Println("Integer Slice:", intSlice)   
	fmt.Println("String Slice:", stringSlice) 
	fmt.Println("Float Slice:", floatSlice) 

	// Length and Capacity of the slices
	fmt.Printf("Length of intSlice: %d, Capacity of intSlice: %d\n", len(intSlice), cap(intSlice))
	// Output: Length of intSlice: 5, Capacity of intSlice: 5

	fmt.Printf("Length of stringSlice: %d, Capacity of stringSlice: %d\n", len(stringSlice), cap(stringSlice))
	// Output: Length of stringSlice: 3, Capacity of stringSlice: 3

	fmt.Printf("Length of floatSlice: %d, Capacity of floatSlice: %d\n", len(floatSlice), cap(floatSlice))
}
