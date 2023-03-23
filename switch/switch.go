package main

import (
	"fmt"
)

// Déclaration d'un tableau contenant les jour de la semaine
var semaines = []string{"lundi", "mardi", "mercrdi", "jeudi", "vendredi", "samedi", "Dimanche"}

func main() {

	for i := 0; i <= 6; i++ { // déclaration d'une boucle pour parcourir les éléments du tableau

		switch i {

		case 0: // Dans le cas ou i est égal à 0
			fmt.Println(semaines[i]) // instruction à exécuter si i = 0 sinon on passe au cas suivant

		case 1:
			fmt.Println(semaines[i])

		case 2:
			fmt.Println(semaines[i])

		case 3:
			fmt.Println(semaines[i])

		case 4:
			fmt.Println(semaines[i])

		case 5:
			fmt.Println(semaines[i])

		case 6:
			fmt.Println(semaines[i])

		default:
			fmt.Println("Nous ne pouvons afficher le jour a la position", i, "du tableau") // s'execute en lieu de l'instruction si la condition d'exécution est pas valide

		}
	}

}
