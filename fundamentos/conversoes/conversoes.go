package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) //preciso converter para float ou para int.
	//ou int(x) / y

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...

	fmt.Println("teste " + string(123))

	fmt.Println("teste " + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
}
