package main

import (
	"fmt"
	"strings"
)

// ModInverse from previous solution
func modInverse(a, m int) int {
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return -1
}

func matMult(a, b, c, d int, pair string) string {
	x := int(pair[0] - 'a')
	y := int(pair[1] - 'a')
	encX := (a*x + b*y) % 26
	encY := (c*x + d*y) % 26
	return string(rune(encX+'a')) + string(rune(encY+'a'))
}

func main() {
	// Key Matrix: [3 3; 2 5]
	k1, k2, k3, k4 := 3, 3, 2, 5
	det := (k1*k4 - k2*k3) % 26
	if det < 0 {
		det += 26
	}
	invDet := modInverse(det, 26)

	// Inverse Matrix for Decryption
	d1 := (k4 * invDet) % 26
	d2 := (-k2 * invDet) % 26
	d3 := (-k3 * invDet) % 26
	d4 := (k1 * invDet) % 26
	// Adjust negatives
	d2 = (d2 + 26) % 26
	d3 = (d3 + 26) % 26

	for {
		fmt.Println("\n--- Hill Cipher (Key: [3 3; 2 5]) ---")
		fmt.Println("1. Encrypt\n2. Decrypt\n3. Exit")
		var ch int
		fmt.Scan(&ch)
		if ch == 3 {
			break
		}

		var text string
		fmt.Print("Enter text (even length): ")
		fmt.Scan(&text)
		text = strings.ToLower(text)
		res := ""

		for i := 0; i < len(text); i += 2 {
			chunk := text[i : i+2]
			if ch == 1 {
				res += matMult(k1, k2, k3, k4, chunk)
			} else {
				res += matMult(d1, d2, d3, d4, chunk)
			}
		}
		fmt.Println("Result:", res)
	}
}
