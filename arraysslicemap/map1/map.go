package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[98765432] = "Pedro"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(cpf, nome)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 12345678) // pra excluir um valor do map
	fmt.Println(aprovados)
}
