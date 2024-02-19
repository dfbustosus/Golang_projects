package main
import "fmt"

// Function to calculate the area of a rectangle
func calculateArea(length, width float64) float64 {
    return length * width
}

func main() {
    // Variable Declaration and Initialization (4 forms)
	// 1 Form
    var name string = "John"
	// 2 Form
    var age int
    age = 30

    //3 Form Type Inference
    var score = 100

    // 4 Form Short Variable Declaration
    height := 6.5

    // Constants
    const pi = 3.14159

    // Print variable values
    fmt.Println("Name:", name)
	fmt.Printf("Type of Name: %T\n", name)
    fmt.Println("Age:", age)
	fmt.Printf("Type of Age: %T\n", age)
    fmt.Println("Score:", score)
	fmt.Printf("Type of Score: %T\n", score)
    fmt.Println("Height:", height)
	fmt.Printf("Type of Height: %T\n", height)
    fmt.Println("Pi:", pi)
	fmt.Printf("Type of Pi: %T\n", pi)

    // Calculate area using the calculateArea function
    length := 5.2
    width := 3.0
    area := calculateArea(length, width)
    fmt.Printf("Area of rectangle with length %.2f and width %.2f is %.2f\n", length, width, area)

    // Zero Value
    var weight int
    fmt.Println("Weight:", weight) // Output: 0 (zero value for int)

    // Scope (Visibility of the variables everyhting outside the {} cannot be used )
    {
        var city string = "New York"
        fmt.Println("City:", city) // Output: New York
    }
    // fmt.Println("City:", city) // Uncommenting this line will cause a compilation error as city is out of scope

}
