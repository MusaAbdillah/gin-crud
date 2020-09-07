package models

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDateBase() {
	database, err := gorm.Open("sqlite3", "/home/musa/gin-crud-test.db")

	if err != nil {
		panic("Failed to connect database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}