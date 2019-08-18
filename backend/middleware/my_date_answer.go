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

type MyDateAnswer struct {
	db *gorm.DB
}

func NewMyDateAnswer(db *gorm.DB) *MyDateAnswer {
	return &MyDateAnswer{
		db: db,
	}
}

func (myDateAnswer *MyDateAnswer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("my date answer")

		vars := mux.Vars(r)
		// そのユーザーがそのハウスにAdminとして所属しているかを確認する
		queryDateAnswerId, ok := vars["questionnaire_date_answer_id"]
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		answerId, err := strconv.ParseInt(queryDateAnswerId, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		user, err := httputil.GetUserFromContext(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		
		isMyAnswer := repository.HEQDAnswerIsMine(myDateAnswer.db, uint(answerId), uint(user.ID))
		if !isMyAnswer {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
