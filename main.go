package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
)

func main() {
	// plainText := "Hola este es lo que tiene que ser secreto"
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error al generar la clave:", err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error al crear el cifrador de bloques AES:", err)
		return
	}

}
