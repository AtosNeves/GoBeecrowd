package main

import (
	"fmt"  // Importa o pacote fmt para operações de entrada e saída.
	"math" // Importa o pacote math para funções matemáticas.
)

func main() {
	var x1, y1, x2, y2 float64 // Declara quatro variáveis do tipo float64 para armazenar as coordenadas dos dois pontos.

	fmt.Scan(&x1, &y1, &x2, &y2) // Lê as coordenadas dos dois pontos da entrada padrão.

	xx := x2 - x1 // Calcula a diferença entre as coordenadas x dos dois pontos.
	yy := y2 - y1 // Calcula a diferença entre as coordenadas y dos dois pontos.

	d := math.Sqrt(math.Pow(xx, 2) + math.Pow(yy, 2)) // Calcula a distância entre os dois pontos usando a fórmula da distância euclidiana.

	fmt.Printf("%.4f\n", d) // Imprime a distância com 4 casas decimais.
}
