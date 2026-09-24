package main 

import "fmt"

func main() {
	flag:= false
	var user string

	if flag {
	user = "admin"
	} else {
	user = "invitado"
	}						

 edad:=0

 fmt.Println("ingrese su edad")	
 fmt.Scan(&edad)

	if edad >= 18 && user == "admin" {
		fmt.Println("Bienvenido\n Tienes todos los privilegios")
	} else if edad < 18 && user == "invitado" {
		fmt.Println("Bienvenido\n Tu usuario no tiene permisos de administrador")

	}else {
		fmt.Println("No puedes ingresar")
	}
}