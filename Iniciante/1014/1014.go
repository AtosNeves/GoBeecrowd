package main

import "fmt"

func main() {
	var distancia float64
	var combustivel float64

	fmt.Scan(&distancia, &combustivel)

	consumo := distancia / combustivel

	fmt.Printf("%.3f km/l\n", consumo)
}
