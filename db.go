package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func init() {
	var err error
	DB, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/gogular")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
}
