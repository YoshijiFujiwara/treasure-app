package model

type User struct {
	BaseModel
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Gender   string   `json:"gender"`
	Birthday string   `json:"birthday"`
	Houses   []*House `gorm:"many2many:house_users;"`
}

// todo うーん、今回はandroidの固有のUUIDでログインできるようにする必要がありそうだ
type RequestLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

// メールアドレス以外をアップデート
type RequestUpdateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}
