package main

import (
	"errors"
	"fmt"
)



func main()  {
	salary := 3000
	if salary <150000 {
		err := errors.New("error: el salario ingresado no alcanza el mínimo imponible")
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}