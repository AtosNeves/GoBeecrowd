package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a, b float32

	fmt.Scan(&a)
	fmt.Scan(&b)

	var c float32

	fmt.Scan(&c)
	var res float32 = b * c
	fmt.Println("NUMBER =", a)
	fmt.Println("SALARY = U$", formatarSalario(res))

}

func formatarSalario(res float32) string {

	var resString string = strconv.FormatFloat(float64(res), 'f', 2, 32)
	return resString
}
