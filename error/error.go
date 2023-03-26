package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) { // lire le contenu d'un fichier et gestion des erreur

	dat, err := os.ReadFile(filename)
	if err != nil { // gestion d'erreur lors de l'ouverture du fichier
		return "", err
	}

	// gestion d'erreur si le fichier est vide
	if len(dat) == 0 {
		// return "", errors.New("Empty content") // créer sa propre érreur
		return "", fmt.Errorf("le fichier %v est vide", filename)
	}
	return string(dat), nil // nil car aucune érreur a été trouvé

}

func main() { // gestion de l'obtention du fichier
	dat, err := readFile("test.txt") // appel
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v", err)
		return // le return permet d'arreter l'exécution des instruction
	}

	fmt.Println("Contenu du fichier :")
	fmt.Println(dat)
}
