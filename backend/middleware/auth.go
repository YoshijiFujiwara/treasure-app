package middleware

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"strings"
)

type Auth struct {
	db     *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		db:     db,
	}
}

func (auth *Auth)Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authorizationヘッダーから、トークンを取得する
		fmt.Println("auth middleware invoked")

		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		authToken := bearerToken[len(bearerToken) - 1]
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("jwt parse error")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println("token parse error")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("token is invalid")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		email := claims["email"].(string)
		// dbからユーザーを検索する
		user := repository.UserFindByEmail(auth.db, email)
		fmt.Println("user")
		fmt.Println(user)
		// ユーザー情報をコンテキストに格納する
		ctx := httputil.SetUserToContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

