package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

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
	router.GET("/productos/:id", getOne)
	router.Run()
}

func getAll(c *gin.Context){
	
	data,err1:= os.ReadFile("productos.json")
	if err1!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err1.Error(),
		})
		return
	}
	var products Products
	err2 := json.Unmarshal(data,&products)
	if err2!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err2.Error(),
		})
		return
	}
	if c.Query("Publicado")!=""{
		b, err4 := strconv.ParseBool(c.Query("Publicado"))
		if err4!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error": err4.Error(),
			})
			return
		}
		var p2 Products
		for _,p :=range products.Productos{
			if p.Publicado==b{
				p2.Productos = append(p2.Productos, p)
			}
		}
		products = p2
	}
	c.JSON(http.StatusOK,products)
}

func getOne(c *gin.Context){
	id := c.Param("id")
	data,err1:= os.ReadFile("productos.json")
	if err1!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err1.Error(),
		})
		return
	}
	var products Products
	err2 := json.Unmarshal(data,&products)
	if err2!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err2.Error(),
		})
		return
	}
	idValue,err3 := strconv.Atoi(id)
	if err2!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err3.Error(),
		})
		return
	}
	product := products.GetById(idValue)
	c.JSON(http.StatusAccepted, product)
}

func (products *Products) GetById(id int) (product Product){
	for _,p:=range products.Productos{
		if p.ID == id{
			product = p
		}
	}
	return
}