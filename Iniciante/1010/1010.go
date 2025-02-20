package main

import "fmt"

func main() {
	var cod1, qt1, cod2, qt2 int
	var valor1, valor2 float64

	fmt.Scan(&cod1, &qt1, &valor1)
	fmt.Scan(&cod2, &qt2, &valor2)

	soma := float64(qt1)*valor1 + float64(qt2)*valor2
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", soma)
}
