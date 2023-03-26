package main

import (
	"fmt"
)

// une slice est un tableau dynamique en go
func main() {

	s := make([]int, 2, 3) // déclaration d'une slice avec trois paramètre dont le tableau, la taille du tableau et la capacité du tableau

	s[0], s[1] = 5, 9 // remplir le tableau

	// affichage de la taille et de la capcité du tableau
	fmt.Println("la taille du tableau est :", len(s), "\nla capacité de mon tableau est:", cap(s)) // len = 2  cap = 2

	s = append(s, 3, 4) // ajouter des éléments au tableau avec pour premier paramètre le tableau et autres comme valeur a y ajouter

	fmt.Println("la taille du tableau est :", len(s), "\nla capacité de mon tableau est:", cap(s)) // len = 4  cap = 6

	fmt.Println(s)

	a := []string{"siaka", "noura", "germain", "mo"} // declaration de tableau sans préciser la taille est une slice

	fmt.Println("la taille du tableau est :", len(a), "\nla capacité de mon tableau est:", cap(a))
	fmt.Println(a)
	fmt.Println("-----------------------------")

	// diviser ma slice
	a1 := a[0:2] // / a1 va contenir a[0] et a[1]

	a2 := a[2:] // divise le tableau a partir de l'index 2 jusque a la fin

	fmt.Println(a1)
	fmt.Println(a2)

	a1[0] = "siahoué" // changer la valeur a l'indice 0

	fmt.Printf("l'élement a la position 1 a été remplacer par %v \n %v", a[0], a)

	fmt.Println("\n-----------------------------")

	new := make([]int, len(s))

	copy(new, s) // copier une slice dans une nouvelle slice

	fmt.Println(len(new), "\n", new) // s a été dupliqué dans new

	// Supprimer un élément d'un tableau

	new = append(new[:0], new[0+1:]...) // supprimer l'élément a l'index 0

	fmt.Println("l'élemnt à l'index 0 a été supprimer : \n", new)

}

// exécution --->  go run slice.go
