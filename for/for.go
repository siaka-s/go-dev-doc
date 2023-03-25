package main

import "fmt"

func main() {

	// déclaration d'une bloucle for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// boucle for avec initialisation de la variable a l'extérieur
	i := 9

	for i >= 0 {
		fmt.Println(i)
		i--
	}

}
