package main

import "fmt"

// defer permet de reporté l'execution d'une instruction a la fermeture de la fonction qui la contient

func start() {

	fmt.Println("start")
}

func finish() {

	fmt.Println("finish")
}

func main() {

	start()
	defer finish() // l'exection en denier lieu dans la fonction qui la contient

	tab := []string{"siaka", "siahoue", "fort"}

	for _, x := range tab {
		fmt.Println(x)
		defer fmt.Println("Good bye", x)
	}

}
