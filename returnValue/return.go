package main

import (
	"fmt"
	"os"
	"strings"
)

// convertir une chaine de caracter en majuscule
func uppercase(word1 string) (upper string, count int) {

	upper = strings.ToUpper(word1) // convertir le mot en majuscule

	count = len(word1) // la fonction len() renvoi la valeur d'un mot

	return upper, count // retourne le mot en majuscule et le nombre de lettre
}

// convertir une chaine de caracter en inuscule
func lowercase(word2 string) (lower string, count int) {

	lower = strings.ToLower(word2) // convertir le mot en minuscule

	count = len(word2)

	return lower, count // retourne le mot en minucule et le nombre de lettre
}

func main() {
	word1 := os.Args[1]

	upper, countupper := uppercase(word1)

	word2 := os.Args[2]

	lower, countlower := lowercase(word2)

	fmt.Printf(" le mot converti en majuscule est %v et sa longueur est %v \n", upper, countupper)

	fmt.Printf(" le mot converti en minuscule est %v et sa longueur est %v", lower, countlower)

}

// exécution ---> go run return.go siahoue SIAKA
