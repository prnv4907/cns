package main

import (
	"bufio"
	"fmt"
	"math/big"
	"net"
	"strings"
)

func main() {
	fmt.Println("1. Server\n2. Client")
	var mode int
	fmt.Scan(&mode)

	// Publicly shared primes
	p := big.NewInt(23) // Prime
	g := big.NewInt(5)  // Generator

	if mode == 1 {
		// --- SERVER ---
		ln, _ := net.Listen("tcp", ":9090")
		fmt.Println("Waiting for client...")
		conn, _ := ln.Accept()

		// 1. Choose Private Key (a)
		a := big.NewInt(6)

		// 2. Calculate Public Key A = g^a mod p
		A := new(big.Int).Exp(g, a, p)

		// 3. Exchange Keys
		// Send A
		fmt.Fprintf(conn, "%s\n", A.String())
		// Receive B
		reader := bufio.NewReader(conn)
		BStr, _ := reader.ReadString('\n')
		B, _ := new(big.Int).SetString(strings.TrimSpace(BStr), 10)

		// 4. Calculate Shared Secret S = B^a mod p
		S := new(big.Int).Exp(B, a, p)
		fmt.Printf("Server Private: %v, Public: %v\nReceived Client Public: %v\nSHARED SECRET: %v\n", a, A, B, S)

	} else {
		// --- CLIENT ---
		conn, _ := net.Dial("tcp", "localhost:9090")

		// 1. Choose Private Key (b)
		b := big.NewInt(15)

		// 2. Calculate Public Key B = g^b mod p
		B := new(big.Int).Exp(g, b, p)

		// 3. Exchange Keys
		// Receive A
		reader := bufio.NewReader(conn)
		AStr, _ := reader.ReadString('\n')
		A, _ := new(big.Int).SetString(strings.TrimSpace(AStr), 10)
		// Send B
		fmt.Fprintf(conn, "%s\n", B.String())

		// 4. Calculate Shared Secret S = A^b mod p
		S := new(big.Int).Exp(A, b, p)
		fmt.Printf("Client Private: %v, Public: %v\nReceived Server Public: %v\nSHARED SECRET: %v\n", b, B, A, S)
	}
}
