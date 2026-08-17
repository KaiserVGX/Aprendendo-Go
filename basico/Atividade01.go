package main

import "fmt"

func main() {
	notas := []float64{
		7.5, 4.0, 8.2, 6.5, 9.0,
		3.5, 5.0, 7.0, 10.0, 6.0,
	}

	aprovados := 0
	recuperacao := 0
	reprovados := 0

	soma := 0.0
	maior := notas[0]
	menor := notas[0]

	for _, nota := range notas {

		if nota >= 7 {
			aprovados++
		} else if nota >= 5 {
			recuperacao++
		} else {
			reprovados++
		}

		soma += nota

		if nota > maior {
			maior = nota
		}

		if nota < menor {
			menor = nota
		}
	}

	media := soma / float64(len(notas))

	fmt.Println("Aprovados:", aprovados)
	fmt.Println("Recuperação:", recuperacao)
	fmt.Println("Reprovados:", reprovados)
	fmt.Println("Maior nota:", maior)
	fmt.Println("Menor nota:", menor)
	fmt.Println("Média:", media)
}
