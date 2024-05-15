package middleware

import (
	"fmt"
	"net/http"

	"github.com/Termpao/auth"
	"github.com/Termpao/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type middlewareRepo struct {
	repo repository.CustomerRepository
}

func NewMiddleAuth(repo repository.CustomerRepository) middlewareRepo {
	return middlewareRepo{repo: repo}
}

func (r *middlewareRepo) Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		done := auth.ParseToken(auth.TokenRequest{
			TokenUser: &token,
		})
		if err != nil || !done {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func (r *middlewareRepo) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenUser, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenUser, func(t *jwt.Token) (interface{}, error) {
			return []byte("suckmydick"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			customer, err := r.repo.GetUser(repository.Customer{Email: claims["email"].(string)})
			if err != nil || customer.Admin == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
			return
		}

		fmt.Println("you not have premission admin")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
