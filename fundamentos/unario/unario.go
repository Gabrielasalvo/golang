package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 5 {
		return "Aprovado"
	}
	return `Reprovado`

}

func main() {
	fmt.Println(obterResultado(5))
}
