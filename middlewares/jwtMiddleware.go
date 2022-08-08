package middlewares

import (
	"fmt"
	"myexample/go-gin/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["athorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(5 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	fmt.Println(config.SecretKey())
	return token.SignedString(config.SecretKey())
}
