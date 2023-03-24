package main

import (
	"fmt"
	"os"
)

func main() {

	pi := 3.14

	nombre := 15

	fmt.Printf("pi est de type %T\n", pi) // affichage du type de la variable

	fmt.Printf("pi a eté converti en %T\n", int(pi)) // Conversion de pi en int et affichage de son nouveau type

	fmt.Printf("nombre est de type %T\n", nombre) // affichage du type de la variable

	fmt.Printf("nombre a été converti en %T\n", float32(nombre)) // Conversion de nombre en float32 et affichage de son nouveau type

	saisi, err := os.ReadFile("file.txt") // lecture et récupération du contenu du fichier file.txt dans une variable saisi et de l'error dans err

	if err != nil { // On vérifie si l'erreur existe
		fmt.Println(err) // dans le cas ou c'est oui on affiche l'error
	}

	fmt.Println(string(saisi)) // Convertion du tableau de byte en chaine de caratère pour l'afficher

}
