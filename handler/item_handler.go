package handler

import (
	"fmt"
	"net/http"

	"github.com/Termpao/service"
	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	service service.ItemService
}

func NewitemHandler(service service.ItemService) itemHandler {
	return itemHandler{service: service}
}

func (h *itemHandler) CreateItem(c *gin.Context) {

	var item service.ItemRes

	if err := c.BindJSON(&item); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	err := h.service.CreateItem(item)

	if err != nil {
		fmt.Println("Cant new product")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)

}

func (h *itemHandler) RemoveItem(c *gin.Context) {
	var items service.ItemRes

	if err := c.BindJSON(&items); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if err := h.service.DeleteItem(int(items.ItemID)); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)
}

func (h *itemHandler) EditItem(c *gin.Context) {
	var item service.ItemRes

	if err := c.BindJSON(&item); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if err := h.service.EditItem(int(item.ItemID), item); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)
}

func (h *itemHandler) GetAllItem(c *gin.Context) {
	items, err := h.service.GetItems()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, items)
}

func (h *itemHandler) GetItem(c *gin.Context) {
	var items service.ItemRes

	if err := c.BindJSON(&items); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	item, err := h.service.GetItem(int(items.ItemID))

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, item)
}
