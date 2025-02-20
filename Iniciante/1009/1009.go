package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var nome string
	fmt.Scan(&nome)

	var salario float32
	fmt.Scan(&salario)
	var vendas float32
	fmt.Scan(&vendas)

	//var total = arredondarCasaDecimal(vendas, salario)
	var stringResultado = formatadorDeSalario(salario, vendas)

	fmt.Println("TOTAL = R$", stringResultado)

}

func formatadorDeSalario(salario float32, vendas float32) string {

	var valor float64 = float64(vendas*0.15 + salario)

	var scale = math.Pow(10, float64(2))
	var res = math.Ceil(valor*scale) / scale

	return strconv.FormatFloat(res, 'f', 2, 32)

}
