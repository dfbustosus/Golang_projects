package main
import "fmt"
// Boolean Constants
const (
	EnableDebugMode = true
	EnableCache     = false
)
// Rune Constants (represents a Unicode code point)
const (
	RuneA = 'A'
	RuneZ = 'Z'
)
// Integer Constants
const (
	MaxConnections = 100
	DaysInWeek     = 7
)
// Floating-Point Constants
const (
	Pi      = 3.14159
	Gravity = 9.81
)
// Complex Constants
const (
	RealPart      = 2.5
	ImaginaryPart = 3.7
)
// String Constants
const (
	Greeting     = "Hello, world!"
	DatabaseName = "my_database"
)

func main() {
	// Boolean Constants
	if EnableDebugMode {
		fmt.Println("Debug mode enabled")
	}

	if !EnableCache {
		fmt.Println("Cache is disabled")
	}
	// Rune Constants
	fmt.Printf("The ASCII value of %c is %d\n", RuneA, RuneA)
	fmt.Printf("The ASCII value of %c is %d\n", RuneZ, RuneZ)
	// Integer Constants
	fmt.Printf("Maximum allowed connections: %d\n", MaxConnections)
	fmt.Printf("Days in a week: %d\n", DaysInWeek)
	// Floating-Point Constants
	fmt.Printf("The value of Pi is %.2f\n", Pi)
	fmt.Printf("The acceleration due to gravity is %.2f m/s²\n", Gravity)
	// Complex Constants
	complexNum := complex(RealPart, ImaginaryPart)
	fmt.Printf("The complex number is: %f + %fi\n", real(complexNum), imag(complexNum))
	// String Constants
	fmt.Println(Greeting)
	fmt.Println("Connecting to database:", DatabaseName)
}
