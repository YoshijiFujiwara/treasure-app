package middleware

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type HouseAdmin struct {
	db     *gorm.DB
}

func NewHouseAdmin(db *gorm.DB) *HouseAdmin {
	return &HouseAdmin{
		db:     db,
	}
}

func (houseAdmin *HouseAdmin)Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("house admin middleware invoked")

		vars := mux.Vars(r)
		// そのユーザーがそのハウスにAdminとして所属しているかを確認する
		queryHouseId, ok := vars["house_id"]
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		houseId, err := strconv.ParseInt(queryHouseId, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		user, err := httputil.GetUserFromContext(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Println("house admin middleware")
		fmt.Println(user.ID)
		fmt.Println(houseId)

		userService := service.NewUser(houseAdmin.db)
		isAdmin, err := userService.TheUserIsAdmin(user.ID, uint(houseId))
		fmt.Println(isAdmin)
		fmt.Println("is admin")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// 権限ないよ〜〜〜
		if !isAdmin {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
