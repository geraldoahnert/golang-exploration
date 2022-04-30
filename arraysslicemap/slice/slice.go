package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de uma array.
	s2 := a2[1:3] // do elemento 1 até o 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas internamente aponta para o array a2
	fmt.Println(s3)

	// vc pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // slice de um slice
	fmt.Println(s2, s4)
}
