package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [...]int{1, 2, 3, 8, 6, 7} //array
	slice := []int{1, 2, 3}           //slice

	fmt.Println(arr, slice)
	fmt.Println(reflect.TypeOf(slice)) // conferindo tipo da estrutura de dados

	a2 := arr[1:4] //slice // NAO E ARRAY é apenas um trecho do array arr
	//tamanho e um ponteiro de um elemento do array
	fmt.Println(a2)
}
