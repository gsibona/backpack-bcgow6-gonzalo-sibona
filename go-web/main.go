package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct{
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}
type Products struct{
	Productos []Product `json:"productos"`
}

func main(){
	router:= gin.Default()
	router.GET("", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola Gonzalo",
		})
	})
	router.GET("/productos", getAll)
	router.Run()
}

var getAll func(c *gin.Context)= func(c *gin.Context){
	
	data,err1:= os.ReadFile("productos.json")
	if err1!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err1.Error(),
		})
	}
	var products Products
	err2 := json.Unmarshal(data,&products)
	if err2!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err2.Error(),
			"data": string(data),
		})
	}
	if err1 == nil && err2 == nil{
		c.JSON(http.StatusOK,products)
	}
}