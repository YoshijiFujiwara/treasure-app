package httputil

import (
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

func GenerateToken(user *model.User) (*model.JWT, error) {
	var err error
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss": os.Getenv("JWT_ISS"),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	jwt := model.JWT{}
	jwt.Token = tokenString

	return &jwt, nil
}



