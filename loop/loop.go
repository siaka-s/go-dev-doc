package main

import "fmt"

func loop() {
	// for i := 0; i < 10; i++ {
	// 	if i < 5 {
	// 		fmt.Printf("le nombre %v est inférieur à 5 \n", i)
	// 	} else if i == 5 {
	// 		fmt.Printf("le nombre %v est égale à 5 \n", i)
	// 	} else if i > 5 {
	// 		fmt.Printf("le nombre %v est supérieur à 5 \n", i)
	// 	}
	// }

	// Deuxième cas de déclaration de boucle
	i := 0
	for i < 10 {
		if i < 5 {
			fmt.Printf("le nombre %v est inférieur à 5 \n", i)
		} else if i == 5 {
			fmt.Printf("le nombre %v est égale à 5 \n", i)
		} else if i > 5 {
			fmt.Printf("le nombre %v est supérieur à 5 \n", i)
		}
		i++
	}
}

// fonction principale s'exécutant en premier dans notre programme
func main() {

	loop()

}
