package models

type User struct {
	Id       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password_hash"`
	Email    string `json:"email" gorm:"column:email"`
	//Role     uint   `json:"role" gorm:"column:role_id"`
}
