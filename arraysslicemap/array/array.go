package main

import "fmt"

func main() {
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7, 8, 9
	fmt.Println(notas)
}
