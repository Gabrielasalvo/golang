package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3}

	for i, numero := range numeros {
		fmt.Printf("%v) %v)", i+1, numero)
	}
}
