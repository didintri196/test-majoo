package util

import (
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
)

var DBConnection struct {
	Session *gorm.DB
}

func ConnectionDB() {
	// fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	// fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	// fmt.Println("DB_DRIVER:", os.Getenv("DB_DRIVER"))
	// fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	// fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	// fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	//open a db connection
	var err error
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DBConnection.Session, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, " + err.Error())
	}
}

func PanicError(err error) {
	log.Println(err)
}
