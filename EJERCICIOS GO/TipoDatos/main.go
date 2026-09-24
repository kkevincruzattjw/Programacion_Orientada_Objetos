package main

import "fmt"

func main() {
	var edad int = 15
	var temperatura float32 = 36.5
	var activo bool = true
	var mensaje string = "Bienvenido a Go"
	var dato byte = 255
	var dias [5]string = [5]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3}

	fmt.Println("edad", edad, "Tu estado es:", activo)
	fmt.Println("temperatura", temperatura)
	fmt.Println("Activo", activo)
	fmt.Println("Mensaje", mensaje)
	fmt.Println("Dato", dato)
	fmt.Println("Dias", dias)
	fmt.Println("Numeros", numeros)

	// Ejercicio 2

	var i int
	var j float64
	fmt.Println("Ingresa dos valores: ")
	fmt.Scanf("%d %f", &i, &j)
	fmt.Println("Resultados:", (float64(i) * j * j))

}

