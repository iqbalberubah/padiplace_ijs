package database

import (
	"fmt"

	e "padiplace_ijs/entity"
	v "padiplace_ijs/vars"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(v.ConnectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}

	DB.AutoMigrate(&e.Product{})
	DB.AutoMigrate(&e.Peternak{})
}
