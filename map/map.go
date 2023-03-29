package main

import (
	"fmt"
)

func main() {

	// déclaration d'une maps

	// marketprice := map[string]int{
	// 	// association de clé <-> valeur
	// 	"huile": 1800,
	// 	"lait":  1500,
	// 	"riz":   1300,
	// 	"pain":  150,
	// }

	// fmt.Println(marketprice["riz"])

	// affichage avec la boucle for range
	// for k, v := range marketprice {

	// 	fmt.Println(k, ":", v)
	// }
	// exemple avec les mois de l'année
	mois := map[int]string{

		1:  "Janvier",
		2:  "Février",
		3:  "Mars",
		4:  "Avril",
		5:  "Mai",
		6:  "Juin",
		7:  "Juillet",
		8:  "Aout",
		9:  "Septembre",
		10: "Octobre",
		11: "Novembre",
		12: "Décembre",
	}

	for k, v := range mois {

		fmt.Println(k, ":", v) // Affichage clé valeur
	}

}
