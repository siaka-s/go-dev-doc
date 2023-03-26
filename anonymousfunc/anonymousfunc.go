package main

import "fmt"

func main() {

	affichage := func() {

		println("je suis une fonction anonyme") // affichage sans import ,envoyer à disparaitre dans le futur
	}

	affichage() // appel de la fonction anonyme

	x, y := func(x, y int) (int, error) {
		println("fonction anonyme")
		return x, nil
	}(10, 15)

	fmt.Printf("la première valeur de retour est :%v \n \bla deuxième valeur de retour est :%v", x, y)

}

// exécution ---> go run anomymousfunc.go
