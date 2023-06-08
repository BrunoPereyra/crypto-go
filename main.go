package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func main() {
	plaintext := []byte("Este es un mensaje secretooo")

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

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		fmt.Println("Error al generar el vector de inicialización:", err)
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	fmt.Println("Texto cifrado:", ciphertext)
	fmt.Println("Clave:", key)

	mode = cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	fmt.Println("Texto decifrado:", string(decrypted))
}
