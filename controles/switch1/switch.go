package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(7.8))
}
