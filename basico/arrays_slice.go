package main

import (
	"fmt"
	"slices"
)

func main() {

	//numeros := [5]int{10, 20, 30, 40} // arrys 5-1
	//numeros[4] = 50
	//fmt.Println("Tamanho", numeros)
	//fmt.Println("Tamanho", len(numeros))

	frutas := []string{"Goiaba", "Maça", "Abacaxi"} //slice
	frutas = append(frutas, "Pera")
	slices.Sort(frutas)
	fmt.Println(frutas)
}
