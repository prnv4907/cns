package main

import (
	"bufio"
	"fmt"
	"math/big"
	"net"
	"strings"
)

// Simple "Textbook" RSA Logic
func encrypt(m *big.Int, e *big.Int, n *big.Int) *big.Int {
	return new(big.Int).Exp(m, e, n)
}

func decrypt(c *big.Int, d *big.Int, n *big.Int) *big.Int {
	return new(big.Int).Exp(c, d, n)
}

func main() {
	fmt.Println("1. Start Server (Receiver/Decryptor)\n2. Start Client (Sender/Encryptor)")
	var mode int
	fmt.Scan(&mode)

	if mode == 1 {
		// --- SERVER ---
		// 1. Generate Keys (Hardcoded primes for demo simplicity)
		p := big.NewInt(61)
		q := big.NewInt(53)
		n := new(big.Int).Mul(p, q)
		phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
		e := big.NewInt(17)
		d := new(big.Int).ModInverse(e, phi)

		fmt.Printf("Server Started. Public Key (e=%v, n=%v)\n", e, n)
		ln, _ := net.Listen("tcp", ":8080")
		conn, _ := ln.Accept()

		// 2. Send Public Key to Client
		fmt.Fprintf(conn, "%s\n%s\n", e.String(), n.String())

		// 3. Receive Encrypted Message
		reader := bufio.NewReader(conn)
		cipherStr, _ := reader.ReadString('\n')
		cipherInt, _ := new(big.Int).SetString(strings.TrimSpace(cipherStr), 10)

		// 4. Decrypt
		plainInt := decrypt(cipherInt, d, n)
		fmt.Printf("Received Cipher: %v\nDecrypted Msg: %v\n", cipherInt, plainInt)

	} else {
		// --- CLIENT ---
		conn, _ := net.Dial("tcp", "localhost:8080")
		reader := bufio.NewReader(conn)

		// 1. Receive Public Key
		eStr, _ := reader.ReadString('\n')
		nStr, _ := reader.ReadString('\n')
		e, _ := new(big.Int).SetString(strings.TrimSpace(eStr), 10)
		n, _ := new(big.Int).SetString(strings.TrimSpace(nStr), 10)

		fmt.Print("Enter numeric message to encrypt: ")
		var msg int64
		fmt.Scan(&msg)
		m := big.NewInt(msg)

		// 2. Encrypt
		c := encrypt(m, e, n)
		fmt.Fprintf(conn, "%s\n", c.String())
		fmt.Printf("Sent Encrypted Message: %v\n", c)
	}
}
