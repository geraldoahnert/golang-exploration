package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":     11324.34,
		"Gabriel Silva": 11324.24,
		"Felipe Felipe": 11324.1,
	}

	funcsESalarios["Felipe Felipe"] = 1113.5

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
