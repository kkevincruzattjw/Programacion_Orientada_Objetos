package main

import (
    "fmt"
	"practica/saludo"
	"practica/operaciones"
	"practica/Taller"
)

func main() {
	fmt.Println("-----------bienvenidos a la clase de paquetes------------")
	mensaje :=saludo.Saludar("KEVIN")
	fmt.Println(mensaje)
	fmt.Println(operaciones.Sumar (5, 6))
	
	var dolares float64
	var moneda string

	fmt.Println("=== CONVERSOR DE MONEDAS ===")
	fmt.Print("Ingrese la cantidad en dólares: ")
	fmt.Scan(&dolares)

	fmt.Print("Ingrese la moneda (Euros, LB, Won, BTC): ")
	fmt.Scan(&moneda)

	Taller.ConvertirMoneda(dolares, moneda)


	var palabra string

	fmt.Println("\n=== CONTADOR DE VOCALES ===")
	fmt.Print("Ingrese una palabra corta: ")
	fmt.Scan(&palabra)

	Taller.ContarVocales(palabra)
}