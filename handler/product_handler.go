package handler

import (
	"fmt"
	"net/http"

	"github.com/Termpao/service"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) productHandler {
	return productHandler{service: service}
}

func (h *productHandler) CreateProduct(c *gin.Context) {

	var product service.ProductReqAndRes

	if err := c.BindJSON(&product); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	product_res, err := h.service.NewProduct(product)

	if err != nil {
		fmt.Println("Cant new product")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, product_res)

}

func (h *productHandler) RemoveProduct(c *gin.Context) {
	var products service.ProductReqAndRes

	if err := c.BindJSON(&products); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if err := h.service.DeleteProduct(products.ProductID); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)
}

func (h *productHandler) EditProduct(c *gin.Context) {
	var product service.ProductReqAndRes

	if err := c.BindJSON(&product); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if err := h.service.EditProduct(product.ProductID, product); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)
}

func (h *productHandler) GetAllProduct(c *gin.Context) {
	products, err := h.service.GetProducts()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, products)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	var product service.ProductReqAndRes

	if err := c.BindJSON(&product); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	products, err := h.service.GetProduct(product.ProductID)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, products)
}
