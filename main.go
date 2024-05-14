package main

import (
	"net/http"

	"github.com/Termpao/handler"
	"github.com/Termpao/middleware"
	"github.com/Termpao/repository"
	"github.com/Termpao/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	dsn := "host=localhost user=postgres password=NADERkungO15 dbname=Termpao port=5432 "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Customer{})

	customerDatabase := repository.NewCustomerDatabase(db)
	customerService := service.NewCustomerService(customerDatabase)

	customerHandler := handler.NewCustomerHandler(customerService)

	// middleware service
	middlewareService := middleware.NewMiddleAuth(customerDatabase)

	router.POST("/login", customerHandler.Login)
	router.POST("/register", customerHandler.Register)

	authorized := router.Group("/")

	authorized.Use(middlewareService.Authentication())
	{
		authorized.POST("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusAccepted, "Heelow")
		})

		admin := authorized.Group("admin")

		admin.Use(middlewareService.Authorization())
		{
			admin.POST("/addcost", customerHandler.AddCost)
		}
	}

	router.Run()
}
