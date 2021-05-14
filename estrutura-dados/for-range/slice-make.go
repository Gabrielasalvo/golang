package main

import "fmt"

func main() {

	s := make([]int, 10) //coloco o tipo de slice e a quantidade de índices, nesse caso 10

	s[3] = 34 //atribuindo valor 34 ao índice 3
	fmt.Println(s[3])

	s = make([]int, 10, 20)
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))
	s = append(s, len(s), 20, 45, 6, 7, 8, 9, 0)

	fmt.Println(s)
}
