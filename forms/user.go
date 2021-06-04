package forms

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserID    int    `json:"id_user" gorm:"column:id_user"`
	Email     string `json:"email" gorm:"column:email"`
	Nama      string `json:"nama" gorm:"column:nama"`
	Password  string `json:"password" gorm:"column:password"`
	Address   string `json:"address" gorm:"column:address"`
	NoTelp    string `json:"no_telp" gorm:"column:no_telp"`
	Gender    string `json:"gender" gorm:"column:gender"`
	MercantID int    `json:"id_mercant" gorm:"column:id_mercant"`
	Role      string `json:"role" gorm:"column:role"`
	LastLogin int    `json:"last_login" gorm:"column:last_login"`
}

type Payload struct {
	Id  string `json:"id"`
	Exp int64  `json:"exp"`
	jwt.StandardClaims
}
