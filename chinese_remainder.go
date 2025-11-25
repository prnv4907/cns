package main

import (
	"fmt"
	"os"
)

// ExtendedGCD finds the Greatest Common Divisor of a and b,
// as well as the coefficients x and y such that ax + by = gcd(a, b).
func ExtendedGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x1, y1 := ExtendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return gcd, x, y
}

// ModInverse finds the modular multiplicative inverse of a under modulo m.
// It returns an error if the inverse does not exist.
func ModInverse(a, m int) (int, error) {
	gcd, x, _ := ExtendedGCD(a, m)
	if gcd != 1 {
		return 0, fmt.Errorf("inverse does not exist (gcd(%d, %d) != 1)", a, m)
	}
	// Ensure the result is positive
	return (x%m + m) % m, nil
}

// CRTReconstruct takes a slice of remainders and a slice of moduli,
// and returns the unique number X (modulo M) using the Chinese Remainder Theorem.
func CRTReconstruct(remainders []int, moduli []int) (int, error) {
	M := 1
	for _, m := range moduli {
		M *= m
	}

	result := 0
	for i, m := range moduli {
		rem := remainders[i]

		// Mi = M / mi
		Mi := M / m

		// yi = Mi^(-1) mod mi
		y, err := ModInverse(Mi, m)
		if err != nil {
			return 0, fmt.Errorf("failed to calculate CRT: %v", err)
		}

		// Formula: Sum(rem * Mi * yi)
		term := rem * Mi * y
		result = (result + term) % M
	}

	return result, nil
}

// ToCRT converts a standard integer into its CRT representation (vector of remainders).
func ToCRT(num int, moduli []int) []int {
	vec := make([]int, len(moduli))
	for i, m := range moduli {
		// Handle negative inputs correctly for modular arithmetic
		vec[i] = ((num % m) + m) % m
	}
	return vec
}

// PerformOperation handles the component-wise arithmetic
func PerformOperation(vecA, vecB, moduli []int, op string) ([]int, error) {
	resultVec := make([]int, len(moduli))

	for i, m := range moduli {
		a := vecA[i]
		b := vecB[i]

		switch op {
		case "+":
			resultVec[i] = (a + b) % m
		case "-":
			resultVec[i] = ((a-b)%m + m) % m
		case "*":
			resultVec[i] = (a * b) % m
		case "/":
			// Division is multiplication by inverse: a * b^(-1) mod m
			invB, err := ModInverse(b, m)
			if err != nil {
				return nil, fmt.Errorf("division impossible: %d has no inverse mod %d", b, m)
			}
			resultVec[i] = (a * invB) % m
		}
	}
	return resultVec, nil
}

func main() {
	var n int
	fmt.Println("--- Chinese Remainder Theorem Calculator ---")
	fmt.Print("Enter the number of moduli (basis size): ")
	fmt.Scan(&n)

	moduli := make([]int, n)
	fmt.Println("Enter the moduli (must be pairwise coprime, e.g., 3 5 7):")
	M := 1
	for i := 0; i < n; i++ {
		fmt.Scan(&moduli[i])
		M *= moduli[i]
	}

	fmt.Printf("Total Modulus Range (M): %d\n", M)
	fmt.Println("--------------------------------------------")

	var num1, num2 int
	fmt.Print("Enter first number (A): ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number (B): ")
	fmt.Scan(&num2)

	// Convert inputs to CRT Vector representation
	vecA := ToCRT(num1, moduli)
	vecB := ToCRT(num2, moduli)

	fmt.Printf("\nCRT Representation of A (%d): %v\n", num1, vecA)
	fmt.Printf("CRT Representation of B (%d): %v\n", num2, vecB)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Add (A + B)")
		fmt.Println("2. Subtract (A - B)")
		fmt.Println("3. Multiply (A * B)")
		fmt.Println("4. Divide (A / B) [Modular Inverse]")
		fmt.Println("5. Change Inputs")
		fmt.Println("6. Exit")
		fmt.Print("Select operation: ")

		var choice int
		fmt.Scan(&choice)

		var resVec []int
		var err error
		var opSymbol string

		switch choice {
		case 1:
			opSymbol = "+"
			resVec, err = PerformOperation(vecA, vecB, moduli, "+")
		case 2:
			opSymbol = "-"
			resVec, err = PerformOperation(vecA, vecB, moduli, "-")
		case 3:
			opSymbol = "*"
			resVec, err = PerformOperation(vecA, vecB, moduli, "*")
		case 4:
			opSymbol = "/"
			resVec, err = PerformOperation(vecA, vecB, moduli, "/")
		case 5:
			fmt.Print("\nEnter new first number (A): ")
			fmt.Scan(&num1)
			fmt.Print("Enter new second number (B): ")
			fmt.Scan(&num2)
			vecA = ToCRT(num1, moduli)
			vecB = ToCRT(num2, moduli)
			fmt.Printf("New A: %v, New B: %v\n", vecA, vecB)
			continue
		case 6:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		if err != nil {
			fmt.Printf("Error performing operation: %v\n", err)
		} else {
			// Reconstruct the result from the vector back to an integer
			finalVal, _ := CRTReconstruct(resVec, moduli)

			fmt.Println("\n--- Result ---")
			fmt.Printf("Vector Operation: %v %s %v = %v\n", vecA, opSymbol, vecB, resVec)
			fmt.Printf("Reconstructed Value (mod %d): %d\n", M, finalVal)

			// Verification for standard operations
			if choice != 4 {
				var actual int
				if choice == 1 {
					actual = num1 + num2
				}
				if choice == 2 {
					actual = num1 - num2
				}
				if choice == 3 {
					actual = num1 * num2
				}

				// Adjust actual to be within modular range for comparison
				actualMod := ((actual % M) + M) % M
				fmt.Printf("Standard Calculation Check: (%d %s %d) %% %d = %d\n", num1, opSymbol, num2, M, actualMod)
			}
		}
	}
}
