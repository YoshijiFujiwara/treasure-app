package service

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/httputil"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/YoshijiFujiwara/u22/backend/repository"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (u *User) Login(loginUser *model.User) (*model.JWT, error) {
	// メールアドレスからユーザーを検索
	user := repository.UserFindByEmail(u.db, loginUser.Email)

	fmt.Println("password")
	fmt.Println(loginUser.Password)
	fmt.Println(user)
	fmt.Println(user.Password)
	// パスワードを比較
	dbPassword := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginUser.Password))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// アクセストークンを発行する
	jwt, err := httputil.GenerateToken(user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return jwt, nil
}

// 新規ユーザーの登録
func (u *User) Create(createUser *model.User) (*model.User, error) {
	user := &model.User{}

	// パスワードハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 10)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user.Username = createUser.Username
	user.Birthday = createUser.Birthday
	user.Gender = createUser.Gender
	user.Email = createUser.Email
	user.Password = string(hash)

	newUser :=repository.UserCreate(u.db, user)
	return newUser, nil
}

func (u *User) Update(id uint, updateUser *model.User) (*model.User, error) {
	user := &model.User{}

	// todo validation

	// パスワードハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), 10)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user.Username = updateUser.Username
	user.Birthday = updateUser.Birthday
	user.Gender = updateUser.Gender
	user.Email = updateUser.Email
	user.Password = string(hash)

	updatedUser := repository.UserUpdate(u.db, id, user)
	return updatedUser, nil
}

func (u *User) Delete(id uint) error {
	repository.UserDelete(u.db, id)
	return nil
}

func (u *User) TheUserIsAdmin(userId uint, houseId uint) (bool, error){
	houseUser := repository.HouseUserFind(u.db, userId, houseId)
	fmt.Println(houseUser)
	fmt.Println("houseuseruseruseurwurewru")
	return houseUser.IsAdmin, nil
}

func (u *User) TheUserBelongsToHouse(userId uint, houseId uint) (bool, error) {
	isExists := repository.HouseUserIsExists(u.db, userId, houseId)
	return isExists, nil
}