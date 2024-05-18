package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenRequest struct {
	TokenUser *string
	EmailUser string
}

func NewToken(data TokenRequest) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.EmailUser,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenstring, err := token.SignedString([]byte("suckmydick"))

	if err != nil {
		return err
	}

	*data.TokenUser = tokenstring
	return nil
}

func ParseToken(data TokenRequest) bool {
	token, err := jwt.Parse(*data.TokenUser, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("suckmydick"), nil
	})
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return false
		}
		// fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
		return false
	}

	return true

}

func ParseTokenData(data TokenRequest) jwt.MapClaims {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(*data.TokenUser, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("suckmydick"), nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	}

	return nil
}
