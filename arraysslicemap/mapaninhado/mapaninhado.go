package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 1234.232,
			"Pedro Silva":   3232.232,
		},
		"J": {
			"João Gui":   23232.3232,
			"Jorge Asde": 29232.2,
		},
	}

	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)

		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
