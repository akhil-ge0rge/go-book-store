package models

import (
	"go_book_store/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func InitBookModel() {
	db = config.GetDB()

	if db == nil {
		panic("database not initialized")
	}

	if err := db.AutoMigrate(&Book{}); err != nil {
		panic(err)
	}
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(b)

	if result.Error != nil {
		return nil, result.Error
	}

	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book

	result := db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func GetBookById(id int64) (*Book, error) {
	var book Book

	result := db.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func DeleteBook(id int64) (*Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil

}

func UpdateBook(id int64, update Book) (*Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	book.Name = update.Name
	book.Author = update.Author
	book.Publication = update.Publication

	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
