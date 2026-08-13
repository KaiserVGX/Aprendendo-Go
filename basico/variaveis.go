package main

import (
	f "fmt"
	m "math"
)

func main() {

	const PI = 3.14
	var raio = 3.2

	f.Print("Informe o valor do raio: ")
	f.Scan(&raio)

	area := PI * m.Pow(raio, 2)
	f.Printf("Àrea é de %.2f\n", area)

}
