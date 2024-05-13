package main

import (
	"github.com/Termpao/handler"
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

	router.POST("/login", customerHandler.Login)
	router.POST("/register", customerHandler.Register)

	router.Run()
}
