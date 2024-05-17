package handler

import (
	"net/http"

	"github.com/Termpao/service"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) orderHandler {
	return orderHandler{service: service}
}

func (h *orderHandler) Order(c *gin.Context) {
	order := service.OrderReq{}

	if err := c.BindJSON(&order); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	var (
		msg *service.OrderRes
		err error
	)

	// Process order based on its status
	if order.Done == 1 {
		msg, err = h.service.OrderSuccess(order)
	} else {
		msg, err = h.service.OrderFail(order)
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, msg)
}
