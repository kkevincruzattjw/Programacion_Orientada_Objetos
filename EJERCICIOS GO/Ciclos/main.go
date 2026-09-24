package main	

import "fmt"

func main(){

	fmt.Println("Bienvenido a la clase  de ciclos")
	fmt.Println("Bucle normal")

	for i:=1; i<100; i++{
		fmt.Println(i)
	}
	fmt.Println("Bucle normal")

	for{
		fmt.Println("Bucle infinito")
		break 
		
	}

	for rango:=range [10]int{}{
		fmt.Println(rango)

	}
}
