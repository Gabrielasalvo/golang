package main

import "fmt"

func main() {

	fmt.Print(`Print`) //print coloca na mesma linha.

	fmt.Println(` coloca uma nova linha`)
	fmt.Println(`nova linha`)

	x := 3.14

	xs := fmt.Sprint(x)
	fmt.Println(`o valor de x eh` + xs)

	a := 1
	b := 2

	fmt.Printf("\n%v %v ", a, b)
}
