package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Termpao/auth"
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

	// var token string
	// if err := auth.NewToken(&token); err != nil {
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	// c.SetCookie("token", string(token), 3600, "/", "localhost", false, true)

	c.Status(http.StatusAccepted)
}

func (s *customerHandler) Login(c *gin.Context) {

	customer := service.CustomerRequest{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}
	if done := auth.ParseToken(&token); !done {
		fmt.Println("asdasd")
		return
	}
}
