package main

import (
	"fmt"
	"os"
)


func Clients() (clients string){
	file, err:= os.ReadFile("./customers.txt")
	defer func()  {
		fmt.Println("ejecución finalizada")
	}()
	if err!=nil{
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return string(file)
}

func main(){
	Clients()
}