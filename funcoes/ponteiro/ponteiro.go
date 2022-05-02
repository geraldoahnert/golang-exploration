package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * nesse caso é para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n)        // por valor, por uma cópia
	fmt.Println(n) // 1

	// revisão: & usado para obter o endereço da variável
	inc2(&n)       // por referência
	fmt.Println(n) // 2
}
