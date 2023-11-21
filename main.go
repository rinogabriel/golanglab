// main.go
package main

import (
	"FunctionAndPackageLab/mathops"
	"fmt"
)

func main() {
	// Test the functions from the mathops package
	result := mathops.Add(10, 5)
	fmt.Println("Addition:", result)

	result = mathops.Subtract(10, 5)
	fmt.Println("Subtraction:", result)

	result = mathops.Multiply(10, 5)
	fmt.Println("Multiplication:", result)

	result = mathops.Power(2, 3)
	fmt.Println("Power:", result)

	quotient, remainder := mathops.Divide(10, 3)
	fmt.Printf("Division: Quotient - %d, Remainder - %d\n", quotient, remainder)
}
