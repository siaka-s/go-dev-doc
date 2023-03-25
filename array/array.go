package main

import "fmt"

func affichage() {

	tab := [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"} // déclaration d'un tableau contenant les jour de la semaine
	i := 0

	for i <= len(tab) {

		fmt.Println("le jour", i+1, "correspond à", tab[i])
		i++
	}

}

func remplissage() {

	var t [10]int // declaration

	i := 0
	// boucle for pour remplir
	for i < 10 {

		t[i] = i
		fmt.Println("je viens de remplir le tableau à la position", t[i])

		i++
	}

}

func main() {
	remplissage()

	fmt.Println("---------------------")

	affichage()
}
