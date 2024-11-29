package Task_1_1

import "fmt"

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

func (lib *Library) RemoveBook(title, release string) {
	for _, book := range lib.Books {
		if book.Title == title && book.ReleaseDate == release {
			delete(lib.Books, title)
		}
	}
}

func (lib *Library) IssueBook() {

}

func (lib *Library) FindBook(title, author string) *Book {

	for _, book := range lib.Books {
		if book.Title == title && book.Author == author {
			return &book
		}
	}

	return nil
}
