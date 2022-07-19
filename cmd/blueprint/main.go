package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/serhijko/go-project-blueprint/cmd/blueprint/apis"
	"github.com/serhijko/go-project-blueprint/cmd/blueprint/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/serhijko/go-project-blueprint/cmd/blueprint/docs"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email serhijko77@gmail.com

// @license.name MIT
// @license.url https://github.com/serhijko/go-project-blueprint/blob/main/LICENSE

// @BasePath /api/v1
func main() {
	// load application configurations
	if err := config.LoadConfig("../../config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", apis.GetUser)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// config.Config.DB.AutoMigrate(&models.User{}) // This is need for generation of schema for postgres image.

	defer config.Config.DB.Close()

	fmt.Println(fmt.Sprintf("Successfully connected to :%v", config.Config.DSN))

	r.RunTLS(fmt.Sprintf(":%v", config.Config.ServerPort), config.Config.CertFile, config.Config.KeyFile)
}
