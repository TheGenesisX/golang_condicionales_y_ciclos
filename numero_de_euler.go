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
	n := 100.0
	i := 3.0
	suma1 := 2.0
	suma2 := suma1 + (1 / factorial(2))

	for i <= n {
		suma1 = suma2
		suma2 += 1 / factorial(i)
		fmt.Println("Suma: ", suma2)
		i++
	}

	fmt.Println("Aproximacion de Euler: ", suma2)
}
