package main

import (
	"fmt"
	"math"
)

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	array1 := array[1:3]

	fmt.Println(array, array1)
	fmt.Printf("Now you have %v problems.\n", math.Sqrt(2))

	// arr := [...]int{1, 2, 3, 8, 6, 7} //array
	// slice := []int{1, 2, 3}           //slice

	// fmt.Println(arr, slice)
	// fmt.Println(reflect.TypeOf(slice)) // conferindo tipo da estrutura de dados

	// a2 := arr[1:4] //slice // NAO E ARRAY é apenas um trecho do array arr
	// //tamanho e um ponteiro de um elemento do array
	// fmt.Println(a2)
}
