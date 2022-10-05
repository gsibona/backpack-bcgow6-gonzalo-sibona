package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/cmd/server/handler"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/internal/products"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/pkg/store"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, os.Getenv("FILEPATH"))
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/",p.Store())
	pr.GET("/",p.GetAll())
	pr.GET("/:id",p.GetById())
	pr.PUT("/:id",p.ModifyAll())
	pr.DELETE("/:id",p.Delete())
	pr.PATCH(":id",p.ModifyValues())
	r.Run()
}