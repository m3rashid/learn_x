package models

import (
	"github.com/jinzhu/gorm"
	"github.com/m3rashid/learn_x/go/bms/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var GetBook Book
	db := db.Where("ID=?", id).Find(&GetBook)
	return &GetBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
