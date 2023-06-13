package jugandounpoco

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func Jugandounpoco(txt string) {

	plaintext := []byte(txt)

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error al generar la clave:", err)
		return
	}

	block, err := aes.NewCipher(key)

	blockSize := block.BlockSize()
	padding := blockSize - (len(plaintext) % blockSize)
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, paddingText...)

	if err != nil {
		fmt.Println("Error al crear el cifrador de bloques AES:", err)
		return
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		fmt.Println("Error al generar el vector de inicialización:", err)
		return
	}
	// cifrado modo cbc
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
