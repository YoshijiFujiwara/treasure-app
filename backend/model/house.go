package model

type House struct {
	BaseModel
	Name      string       `json:"name"`
	Users     []User       `gorm:"many2many:house_users;" json:"user"`
	HouseUser []HouseUser  `gorm:"foreignkey:house_id" json:"house_user"`
	Events    []HouseEvent `gorm:"foreignkey:house_id" json:"events"`
}

type RequestCreateHouse struct {
	Name string `json:"name"`
}
