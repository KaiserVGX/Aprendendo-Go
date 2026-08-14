package main

import "fmt"

func main() {
	var notas [3]float64

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.2
	fmt.Println(notas[2])

}
