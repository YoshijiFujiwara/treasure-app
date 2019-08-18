package model

type HouseUser struct {
	UserID  uint `json: "user_id"`
	HouseID uint `json: "house_id"`
	IsAdmin bool `json: "is_admin"`
}

type RequestAddHouseUser struct {
	IsAdmin bool `json:"is_admin"`
}

type RequestUpdateHouseUser struct {
	IsAdmin bool `json:"is_admin"`
}
