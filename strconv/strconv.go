package main

import (
	"fmt"
	"strconv"
)

func main() {
	// convertions d'une chaine de caractère en entier ----------------------------------------------------------------------------------
	a := "45"

	b := 15

	//som := a + b impossible car la variable a est une chaine de caractère

	conva, err := strconv.Atoi(a) // conversion de a en int pour pouvoir éffevtuer la somme de a et b

	// gestion en cas d'erreur lors de la convertion
	if err != nil {
		fmt.Println(err)
	}

	som := conva + b // conva et b étant de type int la somme peut etre calculer

	fmt.Println(som) // Affichage de la valeur

	// convertion d'un entier en chaine de caractère-------------------------------------------------------------------------------------

	ch := strconv.Itoa(som) // conversion de som en chaine caractère

	fmt.Printf("la variable som a été converti en %T", ch) // affichage du type de la variable ch

}
