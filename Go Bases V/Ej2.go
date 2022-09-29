package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func NewClient(nombreCompleto string, DNI, telefono int, domicilio string){
	legajo:=Legajo()
	if legajo==0{
		panic("error: no se pudo generar el numero de legajo")
	}
	defer func()  {
		err:=recover()
		fmt.Println(err)	
	}()
	clients:=Clients()
	newClient:=fmt.Sprintf("%d,%s,%d,%d,%s;",legajo,nombreCompleto,DNI,telefono,domicilio)
	if strings.Contains(clients,newClient){
		panic("error: este cliente ya se encuentra en la lista")
	}
	if temp, err2 := Check(legajo,nombreCompleto,DNI,telefono,domicilio); !temp{
		panic(err2.Error())
	}
}

func Legajo() int{
	return rand.Int()
	//return 0
}

func Clients() (clients string){
	file, err:= os.ReadFile("./customers.txt")
	defer func()  {
		err:=recover()
		if err!=nil {
			fmt.Println(err)
		}
	}()
	if err!=nil{
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
	return string(file)
}

func Check(legajo int, nombreCompleto string,DNI, telefono int,direccion string)(res bool, err error){
		if legajo<=0||nombreCompleto==""||DNI<=0||telefono<=0||direccion=="" {
			err = errors.New("error: hay valores en su estado 0")
		} else{
			res = true
		}
	return
}

func main()  {
	defer func(){
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")
	}()
	NewClient("gonzalo sibona",12345678,1234567890,"")
}