package book

import "fmt"

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (f *fileRepository) FindAll() ([]Book, error) {
	var books []Book
	fmt.Println("FindAll")
	return books, nil
}

func (f *fileRepository) FindByID(ID int) (Book, error) {
	var book Book
	fmt.Println("FindByID")
	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create")
	return book, nil
}
