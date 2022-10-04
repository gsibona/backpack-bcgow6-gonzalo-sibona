package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/internal/products"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/cmd/server/handler"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/",p.Store())
	pr.GET("/",p.GetAll())
	pr.GET("/:id",p.GetById())
	r.Run()
}