package middleware

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type MyShopAnswer struct {
	db *gorm.DB
}

func NewMyShopAnswer(db *gorm.DB) *MyShopAnswer {
	return &MyShopAnswer{
		db: db,
	}
}

func (myShopAnswer *MyShopAnswer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("my shop answer")

		vars := mux.Vars(r)
		// そのユーザーがそのハウスにAdminとして所属しているかを確認する
		queryShopAnswerId, ok := vars["questionnaire_shop_answer_id"]
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		answerId, err := strconv.ParseInt(queryShopAnswerId, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		user, err := httputil.GetUserFromContext(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		isMyAnswer := repository.HEQSAnswerIsMine(myShopAnswer.db, uint(answerId), uint(user.ID))
		if !isMyAnswer {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
