package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "pq vc nao fala cmg?", 3)
	// fale("João", "só posso falar com vc dps", 1)

	// go fale("Maria", "pq vc nao fala isso comigo?", 500)
	// go fale("João", "não sei tchau", 500)

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns!", 5)
}
