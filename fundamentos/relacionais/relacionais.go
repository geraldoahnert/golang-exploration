package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Strings: ", "Banana" == "Banana") // Strings true
	fmt.Println("!=: ", 3 !=2 ) 
	fmt.Println(3 < 2)
	fmt.Println(3 > 2)
	fmt.Println(3 >= 2)
	
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1 == d2) // true

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("Pessoas: ", p1 == p2) // true
}