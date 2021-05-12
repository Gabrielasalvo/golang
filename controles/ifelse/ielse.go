package main

import "fmt"

func conceito(nota float64) {
	if nota >= 7 && nota <= 10 {
		fmt.Println("conceito A", nota)
	} else if nota >= 5 && nota <= 6 {
		fmt.Println("conceito b", nota)

	} else if nota >= 3 && nota <= 4 {
		fmt.Println("CONCEITO c")
	}
}
func main() {
	conceito(5)
	conceito(7.8)
}
