package main

import "fmt"

func main()  {
	salary := 3000
	if salary <150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}