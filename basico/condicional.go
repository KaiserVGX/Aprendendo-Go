package main

import "fmt"

func media(nota1, nota2, nota3 float64) float64 {
	return (nota1 + nota2 + nota3) / 3
}

func main() {

	var aluno string
	var nota1, nota2, nota3 float64

	fmt.Print("Informe o nome do aluno: ")
	fmt.Scan(&aluno)

	fmt.Print("Informe a nota 01: ")
	fmt.Scan(&nota1)

	fmt.Print("Informe a nota 02: ")
	fmt.Scan(&nota2)

	fmt.Print("Informe a nota 03: ")
	fmt.Scan(&nota3)

	mediaAluno := media(nota1, nota2, nota3)

	fmt.Println("Aluno:", aluno)
	fmt.Println("Média: ", mediaAluno)
}
