package main

import "fmt"

func main() {
	estoque := []int{10, 0, 3, 25, 7, 1, 0, 15}

	semEstoque := 0
	estoqueCritico := 0
	estoqueBaixo := 0
	estoqueNormal := 0

	for _, quantidade := range estoque {

		if quantidade == 0 {
			semEstoque++
		} else if quantidade >= 1 && quantidade <= 3 {
			estoqueCritico++
		} else if quantidade >= 4 && quantidade <= 10 {
			estoqueBaixo++
		} else {
			estoqueNormal++
		}
	}

	fmt.Println("Total sem estoque:", semEstoque)
	fmt.Println("Total em estoque crítico:", estoqueCritico)
	fmt.Println("Total em estoque baixo:", estoqueBaixo)
	fmt.Println("Total em estoque normal:", estoqueNormal)
}
