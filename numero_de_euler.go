package main

import "fmt"

func factorial(numero float64) float64 {

	if numero == 0 {
		return 1
	} else {
		return numero * factorial(numero-1)
	}
}

func main() {
	var n float64
	var i float64 = 1.0
	var suma float64 = 1.0

	fmt.Scanln(&n)

	for i <= n {
		suma += 1 / factorial(i)
		i++
	}

	fmt.Println(suma)
}
