package main

import (
	"fmt"
	"strings"
)

func vigenere(text, key string, encrypt bool) string {
	res := ""
	keyIndex := 0
	for _, char := range strings.ToLower(text) {
		if char < 'a' || char > 'z' {
			res += string(char)
			continue
		}

		shift := int(strings.ToLower(key)[keyIndex] - 'a')
		if !encrypt {
			shift = -shift
		}

		val := int(char - 'a')
		newVal := (val + shift) % 26
		if newVal < 0 {
			newVal += 26
		}

		res += string(rune(newVal + 'a'))
		keyIndex = (keyIndex + 1) % len(key)
	}
	return res
}

func main() {
	var ch int
	var t, k string
	for {
		fmt.Println("\n--- Polyalphabetic Cipher ---")
		fmt.Println("1. Encrypt\n2. Decrypt\n3. Exit")
		fmt.Scan(&ch)
		if ch == 3 {
			break
		}
		fmt.Print("Enter Text: ")
		fmt.Scan(&t)
		fmt.Print("Enter Key: ")
		fmt.Scan(&k)
		fmt.Println("Result:", vigenere(t, k, ch == 1))
	}
}
