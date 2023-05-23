package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablotdv/pocgolang/data"
	"github.com/pablotdv/pocgolang/docs"
	"github.com/pablotdv/pocgolang/models"
	"github.com/pablotdv/pocgolang/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	var err error
	data.Db, err = gorm.Open(mysql.Open("root:123123@tcp(127.0.0.1:3306)/pocgolang?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar no mysql")
	}
	data.Db.AutoMigrate(&models.Pessoa{})

	router := gin.Default()

	api := router.Group("/api")
	{
		pessoa := api.Group("/pessoas")
		{
			pessoa.GET("", routes.GetPessoas)
			pessoa.POST("", routes.PostPessoa)
			pessoa.POST("/sincronizar", routes.PostSincronizarPessoa)
			pessoa.POST("/sincronizar2", routes.PostSincronizarPessoa2)
			pessoa.POST("/sincronizar3", routes.PostPesssoaSincronizar3)
		}
		usuario := api.Group("/usuarios")
		{
			usuario.GET("", routes.GetUsuarios)
		}
	}

	docs.SwaggerInfo.Title = "Swagger POC Golang"
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
