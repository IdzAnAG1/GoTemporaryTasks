package Task_1_1

// File Implementation.go, this File contains realisation function for action with library
import (
	"errors"
	"fmt"
)

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]Book),
	}
}

func (lib *Library) AddBook(newBook Book) {
	lib.Books[newBook.Title] = newBook
	fmt.Printf("Книга - %s, Написанная - %s Добавленна в библиотеку",
		newBook.Title, newBook.Author)
}

func (lib *Library) RemoveBook(title, author string) error {
	if _, exists := lib.Books[author]; exists && lib.bookExists(title) {
		delete(lib.Books, title)
		fmt.Println("Книга удалена")
		return nil
	}
	return errors.New("Такой книги у нас нет, либо данные введены не верно ")
}

func (lib *Library) IssueBook(title, author string) error {
	for _, book := range lib.Books {
		if book.Title == title && book.Author == author && book.Status == true {
			book.Status = false
			return nil
		}
	}
	return errors.New("Нет таких книг в библиотеке ")
}

func (lib *Library) ReturnBook(title, author string) error {
	for _, book := range lib.Books {
		if book.Title == title && book.Author == author && book.Status == false {
			book.Status = true
			return nil
		}
	}
	return errors.New("Данная книга уже возвращена ")
}

func (lib *Library) ShowAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}

func (lib *Library) FindBook(title, author string) *Book {

	for _, book := range lib.Books {
		if book.Title == title && book.Author == author {
			return &book
		}
	}

	return nil
}
