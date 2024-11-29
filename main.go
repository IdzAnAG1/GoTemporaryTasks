package main

import (
	"SIdekick_01/TaskIMPL/Task_1_1"
	"fmt"
)

var (
	title       string
	author      string
	releaseDate string
	status      bool
)

func main() {

	lib := Task_1_1.NewLibrary()
	//
	var book Task_1_1.Book

	fmt.Printf("Введите название книги : ")
	fmt.Scanln(&title)

	fmt.Printf("Введите автора книги : ")
	fmt.Scanln(&author)

	fmt.Printf("Введите дату выпуска книги : ")
	fmt.Scanln(&releaseDate)

	fmt.Printf("Введите дату статус книги : ")
	fmt.Scanln(&status)

	book = Task_1_1.Book{
		Title:       title,
		Author:      author,
		ReleaseDate: releaseDate,
		Status:      status,
	}

	lib.AddBook(book)

	fmt.Println()

	fmt.Println(lib.FindBook(title, author))
}
