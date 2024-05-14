package handler

import (
	"log"
	"net/http"

	"github.com/Termpao/service"
	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	service service.CustomerServiceRepo
}

func NewCustomerHandler(service service.CustomerServiceRepo) customerHandler {
	return customerHandler{service: service}
}

func (s *customerHandler) Register(c *gin.Context) {
	customer := service.CustomerRequest{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if _, err := s.service.Customer_Create(customer); err != nil {
		log.Fatalln("Register error Use service", err)
		return
	}

	c.Status(http.StatusAccepted)
}

func (s *customerHandler) Login(c *gin.Context) {

	customer := service.CustomerRequest{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	user, err := s.service.Customer_Login(customer)

	if err != nil {

		c.Status(http.StatusNoContent)
		return
	}

	// fmt.Println()
	c.JSON(http.StatusOK, gin.H{
		"UseId":    user.CustomerID,
		"Email":    user.Email,
		"Username": user.Username,
	})
	// c.String(http.StatusOK,
}
