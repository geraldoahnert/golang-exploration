package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // ele procura uma case/expressão que é true por padrão
	case t.Hour() < 12:
		fmt.Println("Bom noite!")
	default:
		fmt.Println("Boa noite!")
	}
}
