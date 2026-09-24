package main

import "fmt"

func main() {
	numero := 10

	puntero := &numero

	fmt.Println("valor", numero)
	fmt.Println("direccion de memoria", &numero)
	fmt.Println("puntero", puntero)
	fmt.Println("valor puntero", *puntero)

}
