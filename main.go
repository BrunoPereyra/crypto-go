package main

import (
	generverififirmadig "CRYPTO-GO/gener_verifi_firmaDig"
	"CRYPTO-GO/jugandounpoco"
	"fmt"
)

func main() {
	fmt.Println("eleji que cosa queres hacer pa")
	fmt.Println("1 - generar y verificar una firma digital")
	fmt.Println("2 - cifrar un texto")
	var quecosa int
	fmt.Scanln(&quecosa)
	if quecosa == 1 {
		generverififirmadig.Generverififirmadig()
	} else if quecosa == 2 {
		fmt.Println("agrega texto a cifrar")
		var textoAcifrar string
		fmt.Scanln(&textoAcifrar)
		jugandounpoco.Jugandounpoco(textoAcifrar)
	} else {
		fmt.Println(
			"es invalido, te di dos opciones, 1 o 2, y me pones", quecosa, "😡😡",
		)
	}
}
