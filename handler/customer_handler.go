package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

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

	var token string
	auth.NewToken(auth.TokenRequest{
		TokenUser: &token,
		EmailUser: user.Email,
	})

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"UseId":    user.CustomerID,
		"Email":    user.Email,
		"Username": user.Username,
	})

}

func (s *customerHandler) AddCost(c *gin.Context) {
	customer := service.CustomerRequest{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	err := s.service.Customer_AddMoney(customer)

	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.Status(http.StatusCreated)

}

func (s *customerHandler) ChangePassword(c *gin.Context) {
	customer := service.CustomerRequest{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	err := s.service.Customer_ChangePassword(customer)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)
}

func (s *customerHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Status(http.StatusOK)
}

func (s *customerHandler) TrueWallet_Payment(c *gin.Context) {
	customer := service.CustomerWalletReq{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	formData := url.Values{}
	formData.Set("phone", "0830205297")
	formData.Set("gift_link", customer.CustomerLink)

	walletRes, err := http.PostForm(
		"https://byshop.me/api/truewallet",
		formData,
	)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer walletRes.Body.Close()

	body, err := ioutil.ReadAll(walletRes.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	topUpRes := service.CustomerWalletRes{}

	if err := json.Unmarshal(body, &topUpRes); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if topUpRes.Status == "error" {
		c.Status(http.StatusNoContent)
		return
	}

	customerCost, _ := strconv.Atoi(topUpRes.Amount)
	customerUpdateData := service.CustomerRequest{
		ID:   customer.CustomerID,
		Cost: customerCost,
	}
	if err := s.service.Customer_AddMoney(customerUpdateData); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusAccepted)

}

func (s *customerHandler) BuyItem(c *gin.Context) {
	customer := service.CustomerItemReq{}

	if err := c.BindJSON(&customer); err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if err := s.service.Customer_BuyItem(customer); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (s *customerHandler) GetHistorys(c *gin.Context) {

}
