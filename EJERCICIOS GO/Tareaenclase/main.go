package main

import "fmt"

func promedioNotas(suma float64, cantidad int) float64 {
	return suma / float64(cantidad)
}

func opcionPromedio() {
	var cantidad int

	fmt.Println("\n--- PROMEDIO DE NOTAS ---")
	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&cantidad)

	suma := 0.0

	for i := 1; i <= cantidad; i++ {
		var nota float64

		fmt.Printf("Ingrese la nota del estudiante %d (0 a 100): ", i)
		fmt.Scan(&nota)

		suma = suma + nota
	}

	promedio := promedioNotas(suma, cantidad)

	fmt.Printf("El promedio del curso es: %.2f\n", promedio)

	if promedio >= 70 {
		fmt.Println("El curso está aprobado.")
	} else {
		fmt.Println("El curso está reprobado.")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Excelente rendimiento")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Buen rendimiento")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Rendimiento satisfactorio")
	default:
		fmt.Println("Necesita mejorar")
	}
}

func sumarNumeros() {
	var n int

	fmt.Println("\n--- SUMA DE NÚMEROS ---")
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&n)

	suma := 0

	for i := 1; i <= n; i++ {
		suma = suma + i
	}

	fmt.Println("La suma de los números del 1 al", n, "es:", suma)
}

func convertirCelsiusAFahrenheit() {
	var celsius float64

	fmt.Println("\n--- CELSIUS A FAHRENHEIT ---")
	fmt.Print("Ingrese la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Printf("La temperatura en Fahrenheit es: %.2f\n", fahrenheit)
}

func convertirFahrenheitACelsius() {
	var fahrenheit float64

	fmt.Println("\n--- FAHRENHEIT A CELSIUS ---")
	fmt.Print("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("La temperatura en Celsius es: %.2f\n", celsius)
}

func main() {
	var opcion int

	for {

		fmt.Println("\n===== MENÚ PRINCIPAL =====")
		fmt.Println("1. Promedio de notas")
		fmt.Println("2. Suma de números del 1 al n")
		fmt.Println("3. Celsius a Fahrenheit")
		fmt.Println("4. Fahrenheit a Celsius")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opción: ")

		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			opcionPromedio()

		case 2:
			sumarNumeros()

		case 3:
			convertirCelsiusAFahrenheit()

		case 4:
			convertirFahrenheitACelsius()

		case 0:
			fmt.Println("Gracias por usar el programa.")

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
