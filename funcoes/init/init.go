package main

import "fmt"

func init() {
	fmt.Println("Inicializando...") // executa antes da main
}

func main() {
	fmt.Println("Main...")
}
