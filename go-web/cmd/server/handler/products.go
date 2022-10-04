package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/internal/products"
)

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

type product struct{
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!="password" {
			ctx.JSON(401,gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		p := c.service.GetAll()

		ctx.JSON(200,p)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!="password" {
			ctx.JSON(401,gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		var req product
		if err:=ctx.ShouldBind(&req); err!=nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p,err:=c.service.Store(req.Nombre,req.Color,req.Precio,req.Stock,req.Codigo,req.Publicado,req.FechaCreacion)
		if err!=nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200,p)
	}
}

func (c *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!="password" {
			ctx.JSON(401,gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		id := ctx.Param("id")
		idValue,err := strconv.Atoi(id)
		if err!=nil{
			ctx.JSON(500,gin.H{
				"error": err.Error(),
			})
		return
		}
		p,err:= c.service.GetById(idValue)
		if err!=nil{
			ctx.JSON(404,gin.H{
				"error": err.Error(),
			})
		return
		}
		ctx.JSON(200,p)
	}
}