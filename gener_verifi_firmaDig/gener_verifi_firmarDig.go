package generverififirmadig

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func generateKeyPair() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func signMessage(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func verifySignature(message []byte, signature []byte, publicKey *rsa.PublicKey) error {
	hash := sha256.Sum256(message)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return err
	}
	return nil
}

func Generverififirmadig() {
	// Generar un par de claves pública y privada
	privateKey, err := generateKeyPair()
	if err != nil {
		fmt.Println("Error al generar las claves:", err)
		return
	}

	// Mensaje a firmar
	message := []byte("Este es un mensaje de prueba")

	// Firmar el mensaje
	signature, err := signMessage(message, privateKey)
	if err != nil {
		fmt.Println("Error al firmar el mensaje:", err)
		return
	}

	// Verificar la firma
	err = verifySignature(message, signature, &privateKey.PublicKey)
	if err != nil {
		fmt.Println("Firma no válida:", err)
		return
	}

	fmt.Println("Firma válida")
}
