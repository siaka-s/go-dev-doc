package main

import (
	"fmt"
	"os"
	"strconv"
)

// déclaration de la fonction somme avec deux praramètre de type string x, y et une valeur de retour de type  int
func somme(x, y string) (sum int) {

	xint, _ := strconv.Atoi(x) // convertion de string en int

	yint, _ := strconv.Atoi(y)

	sum = xint + yint

	return sum // la valeur de retour est le résultant de l'exécution de la fonction
}

func main() {

	a := os.Args[1] // permet de récupérer l'argument en position 2 en ligne de commande

	b := os.Args[2] // permet de récupérer l'argument en position 3 en ligne de commande

	// os.Args[0] désigne le nom du fichier qui est  l'argument en position 1 en ligne de commande

	x := somme(a, b) // récupération de la valeur de retour de fonction dans une variable x

	fmt.Printf("la somme de %v + %v est %v\n", a, b, x) // formatage

}

// exécution ----> go run fonction.go 15 15
