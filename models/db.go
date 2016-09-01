package models

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitDB() {
	// Database
	var err error
	var sqlConnection = "root:root@/wedd-in?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("Couldn't connect to database")
	}

	// Migration
	DB.AutoMigrate(Invitation{})
	DB.AutoMigrate(Confirmation{})
}
