package main

import "fmt"

func main() {

	array1 := [3]int{1, 2, 3}

	var slice1 []int

	fmt.Println(array1, slice1)

	slice1 = append(slice1, 45, 67, 89) // adicionando no slice1 os valores 45,67,89
	fmt.Println(slice1)

	slice2 := make([]int, 3)

	copy(slice2, slice1) //pego tudo do slice1 e coloco no slice2 <= slice2 destino

	fmt.Println(slice2, slice1)
}
