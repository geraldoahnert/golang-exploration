package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40" // parecendo que to herdando
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f.nome, f.velocidadeAtual, f.turboLigado)
	fmt.Println(f)
}
