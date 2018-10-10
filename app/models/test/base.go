package test

import (
	db "github.com/hhxsv5/gin-x/db/mysql"
	"github.com/jinzhu/gorm"
)

func Connection() *gorm.DB {
	d, err := db.Connection("test")
	if err != nil {
		panic(err)
	}
	return d
}
