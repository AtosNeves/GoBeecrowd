package main

import (
	"fmt"  // Importa o pacote fmt para operações de entrada e saída.
	"math" // Importa o pacote math para funções matemáticas.
)

func main() {
	var a, b, c float64 // Declara três variáveis do tipo float64 para armazenar os valores de entrada.

	fmt.Scan(&a, &b, &c) // Lê três valores do tipo float64 da entrada padrão e os armazena nas variáveis a, b e c.

	// a) triangulo retangulo
	tri := (a * c) / 2                   // Calcula a área de um triângulo retângulo com base a e altura c.
	fmt.Printf("TRIANGULO: %.3f\n", tri) // Imprime a área do triângulo com 3 casas decimais.

	// b) area do circulo - pi = 3.14159
	area := 3.14159 * math.Pow(c, 2)    // Calcula a área de um círculo com raio c.
	fmt.Printf("CIRCULO: %.3f\n", area) // Imprime a área do círculo com 3 casas decimais.

	// c) a área do trapézio que tem A e B por bases e C por altura.
	areat := (a + b) * c / 2              // Calcula a área de um trapézio com bases a e b e altura c.
	fmt.Printf("TRAPEZIO: %.3f\n", areat) // Imprime a área do trapézio com 3 casas decimais.

	// d) a área do quadrado que tem lado B.
	q := math.Pow(b, 2)               // Calcula a área de um quadrado com lado b.
	fmt.Printf("QUADRADO: %.3f\n", q) // Imprime a área do quadrado com 3 casas decimais.

	// e) a área do retângulo que tem lados A e B
	r := a * b                         // Calcula a área de um retângulo com lados a e b.
	fmt.Printf("RETANGULO: %.3f\n", r) // Imprime a área do retângulo com 3 casas decimais.
}
