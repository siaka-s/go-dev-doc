package main

import (
	"fmt"
	"os"
	"strconv"
)

// Création d'une fonction de gestion de l'entrée utilisateur en argument de ligne de commande
func give() int {

	age := os.Args[1] // le deuxième argument en ligne de commande

	conv, _ := strconv.Atoi(age) // convertion de la valeur de age obtenu en string en entier

	return conv // valeur de retour de la fonction qui sera utiliser pour les conditions

}

func main() {

	if give() < 18 { // condition 1

		fmt.Println(" l'age saisi est celui d'un mineur")

	} else { // condition 2 exécuté si la condition 1 est pas rempli

		fmt.Println(" l'age saisi est celui d'un majeur")
	}
}

// Commande d'exécution         // go run if.go 12             // go run if.go 19
