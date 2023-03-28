package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func traitement(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}

	return found, res, occ
}

func remplacer(src string, dst string, old string, new string) (occ int, lines []int, err error) {

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, nil, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, nil, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := traitement(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}

		fmt.Fprintln(writer, res)
		lineIdx++
	}

	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "Python"
	occ, lines, err := remplacer("texte.txt", "resultat.txt", old, new)
	if err != nil {
		fmt.Printf("Erreur lors du remplacement du mot : %v\n", err)
		return
	}

	fmt.Printf("Nombres d'occurences  %v: %v\n", old, occ)
	fmt.Printf("Nombres de lignes : %d\n", len(lines))
	fmt.Print("Lignes : [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")

}
