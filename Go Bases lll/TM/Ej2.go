package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bProd, err := os.ReadFile("./Products.csv")
	if err!=nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	products := string(bProd)
	fmt.Println("ID\t\t\tPrecio\tCantidad")
	for _, product := range strings.Split(products, "\n"){
		temp := strings.Split(product, ",")
		fmt.Printf("%s\t\t\t%s\t%s\n", temp[0], temp[1], temp[2])
	}
}