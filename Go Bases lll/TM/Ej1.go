package main

import (
	"fmt"
	"os"
)

type Product struct{
	id int
	price float64
	cant int
}

func writeProduct(products ...Product)(){
	file, err := os.ReadFile("./Products.csv")
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, product := range products{
		data := file
		data = append(data, []byte(fmt.Sprintf("\n%d,%f,%d", product.id, product.price, product.cant))...)
		os.WriteFile("./Products.csv", data, 0644)
	}
	if err!=nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main(){
	product := Product{3,50.0,5}
	writeProduct(product)
}