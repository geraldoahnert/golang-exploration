package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável e atribuindo pro ponteiro

	*p++ // desreferenciando (pegando o valr) e incrementando
	i++

	// Go não tem aritimética de ponteiros.
	// p++
	fmt.Println(p, *p, i, &i) // 0xc0000ba000 3 3 0xc0000ba000
}
