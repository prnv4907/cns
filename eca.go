package main

import (
	"fmt"
	"os"
)

// ExtendedGCD recursively calculates the GCD and the coefficients x, y
// such that: a*x + b*y = gcd(a, b)
// Returns: (gcd, x, y)
func ExtendedGCD(a, b int) (int, int, int) {
	// Base Case: If a is 0, then gcd(0, b) is b, and 0*x + b*1 = b
	if a == 0 {
		return b, 0, 1
	}

	// Recursive Step
	// We solve for (b%a)*x1 + a*y1 = gcd
	d, x1, y1 := ExtendedGCD(b%a, a)

	// Update x and y using results from the recursive step
	// Mathematical derivation:
	// x = y1 - (b/a) * x1
	// y = x1
	x := y1 - (b/a)*x1
	y := x1

	return d, x, y
}

// FindModInverse uses ExtendedGCD to find the modular multiplicative inverse.
// It returns the inverse and a boolean indicating success.
func FindModInverse(a, m int) (int, bool) {
	gcd, x, _ := ExtendedGCD(a, m)

	// Inverse only exists if a and m are coprime (gcd is 1)
	if gcd != 1 {
		return 0, false
	}

	// The result 'x' from ExtendedGCD might be negative.
	// In modular arithmetic, we want the positive canonical representation.
	// Example: -2 mod 5 becomes 3.
	result := (x%m + m) % m
	return result, true
}

func main() {
	for {
		fmt.Println("\n--- Extended Euclidean Algorithm Menu ---")
		fmt.Println("1. GCD & Linear Combination (ax + by = gcd)")
		fmt.Println("2. Modular Multiplicative Inverse (a^-1 mod m)")
		fmt.Println("3. Exit")
		fmt.Print("Select operation: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleLinearCombination()
		case 2:
			handleModInverse()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func handleLinearCombination() {
	var a, b int
	fmt.Println("\n--- Linear Combination Calculator ---")
	fmt.Print("Enter integer a: ")
	fmt.Scan(&a)
	fmt.Print("Enter integer b: ")
	fmt.Scan(&b)

	gcd, x, y := ExtendedGCD(a, b)

	fmt.Println("\nResults:")
	fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd)
	fmt.Printf("Coefficients: x = %d, y = %d\n", x, y)
	fmt.Println("Equation Verification (Bézout's Identity):")
	fmt.Printf("(%d * %d) + (%d * %d) = %d\n", a, x, b, y, (a*x + b*y))
}

func handleModInverse() {
	var a, m int
	fmt.Println("\n--- Modular Inverse Calculator ---")
	fmt.Print("Enter number (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter modulus (m): ")
	fmt.Scan(&m)

	inv, ok := FindModInverse(a, m)

	fmt.Println("\nResults:")
	if ok {
		fmt.Printf("The Modular Inverse of %d mod %d is: %d\n", a, m, inv)
		fmt.Println("Verification:")
		fmt.Printf("(%d * %d) %% %d = %d\n", a, inv, m, (a*inv)%m)
	} else {
		fmt.Printf("Inverse does not exist.\n")
		fmt.Printf("Reason: gcd(%d, %d) != 1. They are not coprime.\n", a, m)
	}
}
