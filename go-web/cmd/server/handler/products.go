package handler

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/internal/products"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/pkg/web"
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
	Publicado     *bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

// ListProducts godoc
// @Summary 	List products
// @Tags 		Products
// @Description get products
// @Produce  	json
// @Param token	header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Router 		/products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!="password" {
			ctx.JSON(401,web.NewResponse(401,nil,"no tiene permisos para realizar la peticion solicitada"))
			return
		}

		p,err := c.service.GetAll()
		if err!=nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200,web.NewResponse(200,p,""))
	}
}

// StoreProducts godoc
// @Summary 	Store products
// @Tags 		Products
// @Description store products
// @Accept  	json
// @Produce  	json
// @Param 		token header string true "token"
// @Param 		product body products.Product true "Product to store"
// @Success 	200 {object} web.Response
// @Failure 	401 {object} web.Response
// @Failure 	404 {object} web.Response
// @Router 		/products [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!=os.Getenv("TOKEN") {
			ctx.JSON(401,web.NewResponse(401,nil,"no tiene permisos para realizar la peticion solicitada"))
			return
		}
		var req product
		if err:=ctx.ShouldBind(&req); err!=nil{
			ctx.JSON(404,web.NewResponse(404,nil,err.Error()))
			return
		}
		p,err:=c.service.Store(req.Nombre,req.Color,req.Precio,req.Stock,req.Codigo,req.Publicado,req.FechaCreacion)
		if err!=nil{
			ctx.JSON(404,web.NewResponse(404,nil,err.Error()))
			return
		}
		ctx.JSON(200,web.NewResponse(200,p,""))
	}
}

// GetProductById godoc
// @Summary		Get product by id
// @Tags		Products
// @Description	get product by id
// @Produce 	json
// @Param 		token header string true "token"
// @Param 		id path int true "product id"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router 		/products/{id} [get]
func (c *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!=os.Getenv("TOKEN") {
			ctx.JSON(401,web.NewResponse(401,nil,"no tiene permisos para realizar la peticion solicitada"))
			return
		}
		id := ctx.Param("id")
		idValue,err := strconv.Atoi(id)
		if err!=nil{
			ctx.JSON(500,web.NewResponse(500,nil,err.Error()))
		return
		}
		p,err:= c.service.GetById(idValue)
		if err!=nil{
			ctx.JSON(404,web.NewResponse(404,nil,err.Error()))
		return
		}
		ctx.JSON(200,web.NewResponse(200,p,""))
	}
}

// ModifyProduct godoc
// @Summary 	Modify product
// @Tags 		Products
// @Description modify product
// @Accept  	json
// @Produce  	json
// @Param 		token header string true "token"
// @Param 		id path int true "product id"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router 		/products/{id} [put]
func (c *Product) ModifyAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!=os.Getenv("TOKEN") {
			ctx.JSON(401,web.NewResponse(401,nil,"no tiene permisos para realizar la peticion solicitada"))
			return
		}
		id := ctx.Param("id")
		idValue,err := strconv.Atoi(id)
		if err!=nil{
			ctx.JSON(500,web.NewResponse(500,nil,err.Error()))
			return
		}
		var req product
		if err:=ctx.ShouldBind(&req); err!=nil{
			ctx.JSON(404,web.NewResponse(404,nil,err.Error()))
			return
		}
		p,err:= c.service.ModifyAll(idValue,req.Nombre,req.Color,req.Precio,req.Stock,req.Codigo,req.Publicado,req.FechaCreacion)
		if err!=nil{
			ctx.JSON(404,web.NewResponse(404,nil,err.Error()))
			return
		}
		ctx.JSON(200,web.NewResponse(200,p,""))
	}
}

// DeleteProduct godoc
// @Summary 	Delete product
// @Tags 		Products
// @Description delete product
// @Produce  	json
// @Param token header string true "token"
// @Param 		id path int true "product id"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router 		/products/{id} [delete]
func (c *Product) Delete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!=os.Getenv("TOKEN") {
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
		p,err:=c.service.Delete(idValue)
		if err!=nil{
			ctx.JSON(404,gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200,p)
	}
}

// ModifyNameAndPriceOfProduct godoc
// @Summary 	Modify name and price of product
// @Tags 		Products
// @Description modify name and price of product
// @Produce  	json
// @Param token header string true "token"
// @Param 		id path int true "product id"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 500 {object} web.Response
// @Router 		/products/{id} [patch]
func (c *Product) ModifyValues() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token:= ctx.GetHeader("token")
		if token!=os.Getenv("TOKEN") {
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
		var req product
		if err:=ctx.ShouldBind(&req); err!=nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p,err:= c.service.ModifyValues(idValue,req.Nombre,req.Precio)
		if err!=nil{
			ctx.JSON(404,gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200,p)
	}
}