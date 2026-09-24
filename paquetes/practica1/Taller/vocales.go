package Taller

import "fmt"

func ContarVocales(palabra string) {
    contadorA := 0
    contadorE := 0
    contadorI := 0
    contadorO := 0
    contadorU := 0

    for i := 0; i < len(palabra); i++ {
        letra := palabra[i]

        if letra == 'a' || letra == 'A' {
            contadorA++
        } else if letra == 'e' || letra == 'E' {
            contadorE++
        } else if letra == 'i' || letra == 'I' {
            contadorI++
        } else if letra == 'o' || letra == 'O' {
            contadorO++
        } else if letra == 'u' || letra == 'U' {
            contadorU++
        }
    }

    fmt.Println("--- Conteo de Vocales ---")
    fmt.Println("A:", contadorA)
    fmt.Println("E:", contadorE)
    fmt.Println("I:", contadorI)
    fmt.Println("O:", contadorO)
    fmt.Println("U:", contadorU)
}