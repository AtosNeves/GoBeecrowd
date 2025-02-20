package main

import (
	"fmt"  // Importa o pacote fmt para operações de entrada e saída.
	"math" // Importa o pacote math para funções matemáticas.
)

func main() {
	var a, b, c int // Declara três variáveis inteiras para armazenar os valores de entrada.

	fmt.Scan(&a, &b, &c) // Lê três valores inteiros da entrada padrão e os armazena nas variáveis a, b e c.

	mab := (a + b + int(math.Abs(float64(a-b)))) / 2 // Calcula o maior entre a e b usando a fórmula matemática.
	// math.Abs retorna um float64, então é convertido para int.
	mabc := (mab + c + int(math.Abs(float64(mab-c)))) / 2 // Calcula o maior entre mab (o maior entre a e b) e c.

	fmt.Printf("%d eh o maior\n", mabc) // Imprime o maior valor entre os três números.
}
