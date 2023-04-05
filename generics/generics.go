package main

import (
	"fmt"
)

func main() {

	//--- les generics permettent d'avoir plusieurs types en passage de paramètre de nos fonctions

	fmt.Println(min[bool](false, true, true)) // paramètre de type bool
	fmt.Println(min[int](1, 2, 5))            // paramètre de type int
	fmt.Println(min[float32](5.1, 4.5, 5.7))  // paramètre de type float32

	//--- l'inférence nous permet de ne pas préciser le type de la variable lors de l'appel de la fonction

	pi := 3.14

	fmt.Println(max(0.7, pi)) // paramètre de type float
	fmt.Println(max(7, 5))    // paramètre de type int

}

func min[T int | bool | float32](x, y, z T) T { // t est un type parameter prenant en charge plusieurs types différents, ici int, bool et float32
	return x
}

func max[F float64 | int](x, y F) F { // F est un type parameter prenant en charge deux type qui sont float64 et int
	return y
}
