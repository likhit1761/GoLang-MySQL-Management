package models

import (
	"github.com/jinzhu/gorm"
	"github.com/likhit/BookManagement/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	// Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getbook Book
	db := db.Where("ID=?",Id).Find(&getbook)
	return &getbook,db
}

func DeleteBook(id int64) Book{
	var book Book
	db.Where("ID=?",id).Delete(book)
	return book
}