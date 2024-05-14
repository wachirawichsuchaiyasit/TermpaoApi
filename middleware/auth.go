package middleware

import (
	"net/http"

	"github.com/Termpao/auth"
	"github.com/gin-gonic/gin"
)

func Authorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		done := auth.ParseToken(&token)
		if err != nil || !done {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
