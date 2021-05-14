package main

import "fmt"

func main() {
	// var aprovados map[int]string MAPAS DEVEM SER INICIALIZADOS
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"

	fmt.Println(aprovados)

}
