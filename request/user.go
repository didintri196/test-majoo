package request

type User struct {
	Email     string `json:"email" gorm:"column:email"`
	Nama      string `json:"nama" gorm:"column:nama"`
	Password  string `json:"password" gorm:"column:password"`
	Address   string `json:"address" gorm:"column:address"`
	NoTelp    string `json:"no_telp" gorm:"column:no_telp"`
	Gender    string `json:"gender" gorm:"column:gender"`
	MercantID int    `json:"id_mercant" gorm:"column:id_mercant"`
	Role      string `json:"role" gorm:"column:role"`
}

type UserUpdate struct {
	Email    string `json:"email" gorm:"column:email"`
	Nama     string `json:"nama" gorm:"column:nama"`
	Password string `json:"password" gorm:"column:password"`
	Address  string `json:"address" gorm:"column:address"`
	NoTelp   string `json:"no_telp" gorm:"column:no_telp"`
	Gender   string `json:"gender" gorm:"column:gender"`
	Role     string `json:"role" gorm:"column:role"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
