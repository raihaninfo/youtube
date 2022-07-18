package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id       int `gorm:"primaryKey;autoIncrement:true;unique"`
	Name     string
	Email    string `gorm:"unique:true"`
	Password string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println("db not connected", err.Error())
	}
	db.AutoMigrate(&User{})

	// user := User{
	// 	Id:       2,
	// 	Name:     "Faisal Ahmed",
	// 	Email:    "faisal@gmail.com",
	// 	Password: "1233333",
	// }

	user:= []User{}
	// db.Create(&user)
	db.Find(&user)
	fmt.Println(user)
}
