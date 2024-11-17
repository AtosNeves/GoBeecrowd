package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c, d float32

	fmt.Scan(&a, &b, &c)
	d = calcularMedia(a, b, c, d)
	var resultado = formatar(d)
	fmt.Println("MEDIA =", resultado)
}

func calcularMedia(a, b, c, d float32) float32 {

	return ((a * 2) + (b * 3) + (c * 5)) / 10

}

func formatar(d float32) string {

	return strconv.FormatFloat(float64(d), 'f', 1, 32)

}
