package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	plainAlphabet  = "abcdefghijklmnopqrstuvwxyz"
	cipherAlphabet = "qwertyuiopasdfghjklzxcvbnm" // Example key
)

func transform(input string, from, to string) string {
	var result strings.Builder
	mapping := make(map[rune]rune)
	for i, r := range from {
		mapping[r] = rune(to[i])
	}

	for _, char := range strings.ToLower(input) {
		if val, exists := mapping[char]; exists {
			result.WriteRune(val)
		} else {
			result.WriteRune(char) // Keep punctuation/spaces
		}
	}
	return result.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n--- Monoalphabetic Cipher ---")
		fmt.Println("1. Encrypt")
		fmt.Println("2. Decrypt")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 3 {
			break
		}

		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()

		if choice == 1 {
			fmt.Println("Ciphertext:", transform(text, plainAlphabet, cipherAlphabet))
		} else if choice == 2 {
			fmt.Println("Plaintext:", transform(text, cipherAlphabet, plainAlphabet))
		}
	}
}
