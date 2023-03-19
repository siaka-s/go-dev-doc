package main

import "fmt"

// déclaration global, explicite avec des valeur de type identiques
var name, surname string = "kouadio", "piere"

// déclaration multiple
var (
	age uint = 15 // uint est un type qui autorise pas les nombres négatifs

	genre string // déclaré mais non initialisé avec une chaine vide par défaut

	taille float32 = 1.86

	naissance = "03/05/2023" // go va déduire le type de la variable
)

func main() {

	genre = "masculin" // initialiasation

	// déclaration implicite autorisée à l'intérieur d'une fonction
	location := "abidjan"

	fmt.Println(name, surname, age, genre, location, taille)
}

// test := "test"
// la ligne du dessus renvera une érreur car la déclaration impicite est pas autorisé ici
