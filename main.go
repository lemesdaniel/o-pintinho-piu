package main

import (
	"fmt"
)

// clássico da música Brasileira
func main() {

	animais := []string{"pintinho", "galinha", "galo", "pelu", "capote", "gato", "cacholo", "cabla", "bode", "vaca", "boi", "moça"}
	sons := []string{"Piu", "Có", "Cocolicó", "Glu Glu", "To Flaco", "Miau", "Au Au", "Mééé", "Bééé", "Muuuu", "Móóó", "Oh!"}

	var animal []string

	for _, v := range animais {

		switch v {
		case "moça", "vaca", "cabra", "galinha":
			frase(2, v, "uma")
		case "boi", "bode", "cachorro", "gato", "capote", "peru", "galo", "pintinho":
			frase(2, v, "um")

		}

		animal = append(animal, v)
		for n := len(animal) - 1; n >= 0; n-- {
			switch animal[n] {
			case "moça", "vaca", "cabra", "galinha":
				frase(2, "a", animal[n], sons[n])

			case "boi", "bode", "cachorro", "gato", "capote", "peru", "galo", "pinto":
				frase(2, "o", animal[n], sons[n])
			case "pintinho":
				frase(4, "o", animal[n], sons[n])

			}
		}

		fmt.Println("---------------------")
	}
}

func frase(n int, texto ...string) {
	if len(texto) == 3 {
		for i := 0; i < n; i++ {
			fmt.Printf("E %s %s %s\n", texto[0], texto[1], texto[2])
		}
	} else {
		for i := 0; i < 2; i++ {
			fmt.Printf("Lá em casa tinha %s %s\n", texto[0], texto[1])
		}
	}

}
