package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/cmd/server/handler"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/internal/products"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/pkg/store"
	"github.com/joho/godotenv"
	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/docs"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Meli Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService http://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, os.Getenv("FILEPATH"))
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	docs.SwaggerInfo.Host=os.Getenv("HOST")
	r.GET("/docs/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.POST("/",p.Store())
	pr.GET("/",p.GetAll())
	pr.GET("/:id",p.GetById())
	pr.PUT("/:id",p.ModifyAll())
	pr.DELETE("/:id",p.Delete())
	pr.PATCH(":id",p.ModifyValues())
	if err := r.Run(); err!=nil{
		panic(err)
	}
}