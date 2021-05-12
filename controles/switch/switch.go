package main

import "fmt"

func notas(n float64) string {

	switch {
	case n <= 10 && n >= 7:
		return ` TIROU DEZZZ`
	default:
		return `TIROU ZERO`
	}

}
func main() {
	fmt.Println(notas(7))
}
