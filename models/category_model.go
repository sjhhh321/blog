package models

type CategoryModel struct {
	Model
	Title     string    `json:"title"`
	UserID    uint      `json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
}
