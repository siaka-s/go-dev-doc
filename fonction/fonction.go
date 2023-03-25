package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(x, y string) (sum int) {

	xint, _ := strconv.Atoi(x)

	yint, _ := strconv.Atoi(y)

	sum = xint + yint

	return sum
}

func main() {

	a := os.Args[1]

	b := os.Args[2]

	x := sum(a, b)

	fmt.Printf("la somme de %v + %v est %v\n", a, b, x)

}
