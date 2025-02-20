package main

import "fmt"

func main() {

	var a, b float32

	fmt.Scan(&a, &b)

	var c = (a*3.5 + b*7.5) / 11

	fmt.Printf("MEDIA = %.5f\n", c)

}
