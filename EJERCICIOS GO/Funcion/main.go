package main

import "fmt"

func calcular(a, b int) (int, int) {
	if a > b {
		return a + b, b - a
		return a + b, a - b
	} else {
		return a + b, 0
	}
}

func main() {
	s, r := calcular(10, 5)
	fmt.Println("Suma:", s)
	fmt.Println("Resta:", r)
}	