package main

import "fmt"

func main(){
	
	//Ej 1

	var nombre = "Gonzalo Sibona"
	var direccion = "Las Achiras 380"

	fmt.Println(nombre)
	fmt.Println(direccion)


	//Ej 2

	var temperatura = 17
	var humedad = 45
	var presion = 1013.8

	fmt.Printf("%d°C ", temperatura)
	fmt.Printf("%d%% ", humedad)
	fmt.Printf("%g hPa ", presion)
	fmt.Println("")

	//Ej 3

	/*
	var nombre string //mal declarado, 1nombre => nombre
	var apellido string 
	var edad int // mal declarado, primero va el nombre de la variable y despues el numero
	apellido := 6 //mal declarado, 1apellido => apellido, 
	var licencia_de_conducir = true
	var estatura_de_la_persona int //mal declarado, estatura de la persona => estatura_de_la_persona
	cantidadDeHijos := 2
	*/


	//Ej 4

	var apellido_2 string = "Gomez"
  	var edad int = 35
  	boolean := false;
  	var sueldo float64 = 45857.90
  	var nombre_2 string = "Julián"

	fmt.Println(apellido_2, edad, boolean, sueldo, nombre_2)

}