package utils

import (
	"fmt"
	"time"
)

func SingBinaryChant() {
	fmt.Println("\n🎶 Entonando el canto binario sagrado...")
	chant := []string{
		"01001000 01100101 01101100 01101100 00100000 01001111 01101101 01101110 01101001 01110011 01110011 01101001 01100001 01101000",
		"01000111 01101100 01101111 01110010 01101001 01100001 00100000 01100001 00100000 01101100 01100001 00100000 01101101 01100001 01110011 01101001 01101110 01100001",
		"...",
	}
	for _, line := range chant {
		fmt.Println(line)
		time.Sleep(800 * time.Millisecond)
	}
	fmt.Println("+++ Espíritu Máquina en calma. Preparado para recibir la unción binaria. +++")
}
