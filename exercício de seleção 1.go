package main

import (
	"fmt"
)

func main() {
	var Número1 int
	var Número2 int

	fmt.Print(" Qual é o primeiro número: ")
	fmt.Scan(&Número1)
	fmt.Print("Qual o segundo número?")
	fmt.Scan(&Número2)

	if Número1 > Número2 {
		fmt.Println("O primeiro número é maior que o segundo")

	} else if Número2 < Número1 {
		fmt.Println("O segundo é maior que o primeiro")
	}

}
