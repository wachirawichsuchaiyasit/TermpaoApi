package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Termpao/handler"
	"github.com/Termpao/middleware"
	"github.com/Termpao/repository"
	"github.com/Termpao/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.AutoMigrate(&repository.Customer{})
	db.AutoMigrate(&repository.Product{})
	db.AutoMigrate(&repository.ItemOrder{})
	db.AutoMigrate(&repository.Order{})

	// repository
	customerDatabase := repository.NewCustomerDatabase(db)
	productDatabase := repository.NewProductRepository(db)
	itemRepository := repository.NewItemRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	// Service
	customerService := service.NewCustomerService(customerDatabase)
	productService := service.NewProductService(productDatabase)
	itemService := service.NewItemService(itemRepository)
	orderService := service.NewOrderService(orderRepository)

	// handler
	customerHandler := handler.NewCustomerHandler(customerService)
	productHandler := handler.NewProductHandler(productService)
	itemHandler := handler.NewitemHandler(itemService)
	orderHandler := handler.NewOrderHandler(orderService)

	// middleware service
	middlewareService := middleware.NewMiddleAuth(customerDatabase)

	router.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/login", customerHandler.Login)
	router.POST("/register", customerHandler.Register)

	authorized := router.Group("/")

	authorized.Use(middlewareService.Authentication())
	{
		// Customer
		authorized.POST("/editpassword", customerHandler.ChangePassword)
		authorized.POST("/logout", customerHandler.Logout)
		authorized.POST("/wallet_topup", customerHandler.TrueWallet_Payment)
		authorized.POST("/buy", customerHandler.BuyItem)
		authorized.GET("/@me", customerHandler.GetData)
		authorized.GET("/items", itemHandler.GetItemsProduct)

		// Admin Premission
		admin := authorized.Group("admin")
		admin.Use(middlewareService.Authorization())
		{
			admin.POST("/addcost", customerHandler.AddCost)

			// Product
			admin.POST("/product", productHandler.CreateProduct)
			admin.DELETE("/product", productHandler.RemoveProduct)
			admin.PUT("/product", productHandler.EditProduct)
			admin.GET("/products", productHandler.GetAllProduct)
			admin.GET("/product", productHandler.GetProduct)

			// Items
			admin.POST("/item", itemHandler.CreateItem)
			admin.DELETE("/item", itemHandler.RemoveItem)
			admin.PUT("/item", itemHandler.EditItem)
			admin.GET("/items", itemHandler.GetAllItem)
			admin.GET("/item", itemHandler.GetItem)

			// Order
			admin.POST("/order", orderHandler.Order)

		}
	}

	router.Run(":" + os.Getenv("PORT"))
}
