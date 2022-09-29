package main

import (
	"fmt"
)

type ErrorSalary struct{

}
func (e *ErrorSalary) Error() string{
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main()  {
	salary := 3000
	if salary <150000 {
		var err ErrorSalary
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}