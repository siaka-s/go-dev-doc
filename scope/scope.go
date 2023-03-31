package main

import (
	"fmt"
)

// déclaration globale
var x int

func main() {
	x = 5
	f() // 10 car x est déclaré comme variable local

	fmt.Println(x)
	T()

}

func f() {
	// variable local déclarée dans la fonction
	x := 10
	fmt.Println(x)

}

func T() {

	fmt.Println(x)
	// x récupére la valeur globale
}
