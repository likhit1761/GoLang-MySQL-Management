package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","root:rootlikhit@tcp(127.0.0.1:3306)/gotable1")
	if err!=nil{
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}