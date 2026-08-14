package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	compraTv50 := trab1 && trab2
	compraTv32 := trab1 != trab2
	comprarSorverte := trab1 || trab2

	return compraTv50, compraTv32, comprarSorverte
}

func main() {
	tv50, tv32, soverte := compras(true, false)

	fmt.Printf("Tv50: %t, Tv32: %t, soverte: %t, Saudável %t\n",
		tv50, tv32, soverte, !soverte)
}
