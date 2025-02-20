package main

import "fmt"

func main() {

	var a, b, c, d, e int

	fmt.Scan(&a, &b, &c, &d)
	e = acharDiferenca(a, b, c, d)
	fmt.Println("DIFERENCA =", e)

}

func acharDiferenca(a, b, c, d int) int {

	var e int = (a * b) - (c * d)

	return e

}
