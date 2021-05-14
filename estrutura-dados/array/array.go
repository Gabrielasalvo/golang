package main

import "fmt"

func main() {
	//estrutura homogenea (mesmo tipo) e estatica (mesmo tamanho sempre)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7, 1.5, 9.1

	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("A media é %f %f", media, media/2)
	console.log(`${media}`)
	console.log("minha nota é {}", media)
}
