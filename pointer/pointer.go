package main

import "fmt"

func maj(name *string) { // name est un pointer de type string pointant vers l'adresse mémoire de firstname

	*name = "siahoué"

}

func main() {
	firstname := "siaka"

	maj(&firstname) // &firstname représente l'adresse mémoire de la variable firstname

	fmt.Println(firstname)

}
