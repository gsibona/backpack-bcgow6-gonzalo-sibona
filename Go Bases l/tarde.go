package main

import "fmt"

func main () {

	//Ej 1
	palabra := "palabra"
	fmt.Println(len(palabra))
	for _,i := range palabra{
		fmt.Println(string(i))
	}

	//Ej 4
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de benjamin es: ", employees["Benjamin"])
	x:= 0
	for _,i := range employees {
		if i>21 {
			x++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 es: ", x)
	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)

}


//Ej 2

func banco(edad int, empleado bool, antiguedad int, sueldo int) {
	if edad <= 22 {
		fmt.Println("Se a denegado acceso al prestamo ya que tiene menos de 23 años")
	} else if !empleado {
		fmt.Println("Se a denegado acceso al prestamo ya que no se encuentra en relacion de dependencia")
	} else if antiguedad <= 1 {
		fmt.Println("Se a denegado acceso al prestamom ya que no cumple con el minimo de 2 años de antiguedad")
	} else if sueldo < 100000 {
		fmt.Println("Prestamo aprobado con intereses")
	} else {
		fmt.Println("Prestamo aprobado sin intereses")
	}
}

//Ej 3
func mes(mes int){
	switch mes {
	case 1: fmt.Println("Enero")
	case 2: fmt.Println("Febrero")
	case 3: fmt.Println("Marzo")
	case 4: fmt.Println("Abril")
	case 5: fmt.Println("Mayo")
	case 6: fmt.Println("Junio")
	case 7: fmt.Println("Julio")
	case 8: fmt.Println("Agosto")
	case 9: fmt.Println("Septiembre")
	case 10: fmt.Println("Octubre")
	case 11: fmt.Println("Noviembre")
	case 12: fmt.Println("Diciembre")
	default: fmt.Println(mes, " no es un numero de mes valido")	
	}
}