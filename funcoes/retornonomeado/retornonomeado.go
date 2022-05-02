package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1

	return // retorno limpo, já declarei as variáveis de retorno então não preciso especificar
}

func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)
}
