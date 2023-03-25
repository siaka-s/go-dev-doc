package main

import (
	"fmt"
	"os"
	"strconv"
)

func affiche() {

	sasisi := os.Args[1] // récuperer le nombre en argument de ligne de commande

	i, _ := strconv.Atoi(sasisi) // convertion string en int

	fmt.Println("le résultat du premier flux de controle avec if/else :")
	// Controle de flux avec if/else

	if i <= 10 {
		if i < 0 {
			fmt.Println("la valeur saisi ne peut pas etre traité ")
		} else if i < 5 && i >= 0 {
			fmt.Println("le nombre", i, "est plus petit que 5")
		} else if i == 5 {
			fmt.Println("le nombre", i, "est égale à 5")
		} else if i > 5 && i >= 0 {
			fmt.Println("le nombre", i, "est supérieur à 5")
		}
	}

	fmt.Println("le résultat du deuxièeme flux de controle avec switch :")
	// Controle de flux avec switch simplifie la lecture du code
	switch i <= 10 {
	case i < 0:
		fmt.Println("la valeur saisi ne peut pas etre traité ")
	case i < 5 && i >= 0:
		fmt.Println("le nombre", i, "est plus petit que 5")
	case i == 5:
		fmt.Println("le nombre", i, "est égale à 5")
	case i > 5 && i >= 0:
		fmt.Println("le nombre", i, "est supérieur à 5")
	}
}

func main() {

	affiche()

}

// exécution --> go run flux.go 5
