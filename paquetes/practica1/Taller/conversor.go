package Taller

import "fmt"

func ConvertirMoneda(dolares float64, moneda string) {
    if moneda == "Euros" {
        fmt.Println("Resultado:", dolares * 0.92, "Euros")
    } else if moneda == "LB" {
        fmt.Println("Resultado:", dolares * 0.79, "Libras")
    } else if moneda == "Won" {
        fmt.Println("Resultado:", dolares * 1330, "Wones")
    } else if moneda == "BTC" {
        fmt.Println("Resultado:", dolares / 65000, "BTC")
    } else {
        fmt.Println("Moneda no permitida o mal escrita.")
    }
}