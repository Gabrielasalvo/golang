package main

import "fmt"

func compras(trab, trab2 bool) (bool, bool, bool) { //colocando retorno dentro de parenteses
	comprarTv := trab && trab2
	comprarTv2 := trab != trab2 //OU exclusivo
	comprarSorvete := trab || trab2

	return comprarTv, comprarTv2, comprarSorvete

}

func main() {

	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf(`tv50: %t, tv32: %t, sorvete: %t, saudavel: %t`, tv50, tv32, sorvete, !sorvete)
}
