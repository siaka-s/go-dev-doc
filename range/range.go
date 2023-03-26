package main

import "fmt"

func tab() {

	tab := [5]int{1, 2, 3, 4, 5} // declaration du tableau

	//syntaxe d'utilisation de la boucle avec range pour afficher des couples index/valeur
	for index, valeur := range tab {

		fmt.Printf("le numero assoyé au rang %v du tableau est %v\n", index, valeur)
	}
}
func main() {

	tab()
}
