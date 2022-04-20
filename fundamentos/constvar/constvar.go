package main

import (
	"fmt"
	m "math" // m as math
)

func main(){
	const PI float64 = 3.1415
	var raio = 3.2 // compilador sabe que é (float 64)

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("Area da circuferencia é", area)

	// Outras formas de declarar
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "uepa!"
	fmt.Println(g, h, i)
}