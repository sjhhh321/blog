package models

type UserModel struct {
	Model
	Username       string `gorm:"size:32" json:"username"`
	Nickname       string `gorm:"size:32" json:"nickname"`
	Avatar         string `gorm:"size:256" json:"avatar"`
	Abstract       string `gorm:"size:256" json:"abstract"`
	RegisterSource int8   `json:"registerSource"`
	CodeAge        int    `json:"codeAge"`
	Password       string `gorm:"size:64" json:"-"`
	Email          string `gorm:"size:256" json:"email"`
	OpenID         string `gorm:"size:64" json:"openID"` //第三方登录的唯一id
}
