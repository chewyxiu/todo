package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	dsn := "root:@/todo?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	db.LogMode(true)

	fmt.Printf("Success: Connected to DB!!\n")
}

