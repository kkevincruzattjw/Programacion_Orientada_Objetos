
package main

import "fmt"

func saludar() { 
	fmt.Println("Hola esta es mi primer funcion")

}

func bienveinido(nombre string) {
	fmt.Println("Bienvenido", nombre)
}

func calcular(a, b int) (int, int) {
	if a > b {
		return a + b, b - a
		return a + b, a - b
	} else {
		return a + b, 0
	}
}

func sumatoria(num int) int {
	total := 0
	for _,num:= range numeros {
		total += num
	}
	return total
}

func main() {
	var usr string 

	fmt.Println("Escribe tu nombre")
	fmt.Scanln(&usr)
	saludar()
	bienveinido(usr)

	s, r := calcular(10, 5)
	fmt.Println("Suma:", s)
	fmt.Println("Resta:", r)
	mostrarNum(5, 10 ,15, 20, 25)
	fmt.Println("La sumatoria es:", sumatoria(1, 2, 3, 4, 5,6 7, 8, 9,))
}
