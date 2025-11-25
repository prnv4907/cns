package main

import (
	"fmt"
	"strings"
)

func prepareKeyMatrix(key string) [5][5]rune {
	key = strings.ReplaceAll(strings.ToLower(key), "j", "i")
	alphabet := "abcdefghiklmnopqrstuvwxyz" // 'j' omitted
	matrixStr := ""
	seen := make(map[rune]bool)

	addChar := func(r rune) {
		if !seen[r] && (r >= 'a' && r <= 'z') {
			matrixStr += string(r)
			seen[r] = true
		}
	}

	for _, r := range key {
		addChar(r)
	}
	for _, r := range alphabet {
		addChar(r)
	}

	var matrix [5][5]rune
	for i, r := range matrixStr {
		matrix[i/5][i%5] = r
	}
	return matrix
}

func findPos(matrix [5][5]rune, char rune) (int, int) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if matrix[r][c] == char {
				return r, c
			}
		}
	}
	return -1, -1
}

func process(text, key string, encrypt bool) string {
	matrix := prepareKeyMatrix(key)
	text = strings.ReplaceAll(strings.ToLower(text), "j", "i")
	text = strings.ReplaceAll(text, " ", "")

	// Pad digraphs
	var digraphs []string
	for i := 0; i < len(text); i += 2 {
		if i+1 == len(text) {
			digraphs = append(digraphs, string(text[i])+"x")
		} else if text[i] == text[i+1] {
			digraphs = append(digraphs, string(text[i])+"x")
			i-- // re-process second char
		} else {
			digraphs = append(digraphs, text[i:i+2])
		}
	}

	shift := 1
	if !encrypt {
		shift = 4
	} // Move back (equivalent to +4 mod 5)

	var result string
	for _, pair := range digraphs {
		r1, c1 := findPos(matrix, rune(pair[0]))
		r2, c2 := findPos(matrix, rune(pair[1]))

		if r1 == r2 { // Same Row
			result += string(matrix[r1][(c1+shift)%5]) + string(matrix[r2][(c2+shift)%5])
		} else if c1 == c2 { // Same Col
			result += string(matrix[(r1+shift)%5][c1]) + string(matrix[(r2+shift)%5][c2])
		} else { // Rectangle
			result += string(matrix[r1][c2]) + string(matrix[r2][c1])
		}
	}
	return result
}

func main() {
	var choice int
	var text, key string
	for {
		fmt.Println("\n--- Playfair Cipher ---")
		fmt.Println("1. Encrypt\n2. Decrypt\n3. Exit")
		fmt.Scan(&choice)
		if choice == 3 {
			break
		}
		fmt.Print("Enter Key: ")
		fmt.Scan(&key)
		fmt.Print("Enter Text: ")
		fmt.Scan(&text)
		if choice == 1 {
			fmt.Println(process(text, key, true))
		}
		if choice == 2 {
			fmt.Println(process(text, key, false))
		}
	}
}
