package database

import (
	"dewetour/models"
	"dewetour/pkg/mysql"
	"fmt"
)

func Migrate() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration success")
}
