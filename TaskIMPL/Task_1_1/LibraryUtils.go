package Task_1_1

import "errors"

func (lib *Library) bookExists(title string) bool {
	if _, exists := lib.Books[title]; exists {
		return true
	}
	return false
}

func (lib *Library) searchBookByAuthor(author string) ([]Book, error) {

	FoundBooks := make([]Book, 0, len(lib.Books))

	for _, book := range lib.Books {
		if book.Author == author {
			FoundBooks = append(FoundBooks, book)
		}
	}
	if len(FoundBooks) == 0 {
		return nil, errors.New("Нет книг с таким автором ")
	}
	return FoundBooks, nil
}

func (lib *Library) searchBookByTitle(title string) ([]Book, error) {

	FoundBooks := make([]Book, 0, len(lib.Books))

	for _, book := range lib.Books {
		if book.Title == title {
			FoundBooks = append(FoundBooks, book)
		}
	}
	if len(FoundBooks) == 0 {
		return nil, errors.New("Нет книг с таким названием ")
	}
	return FoundBooks, nil
}

func (lib *Library) searchBookByRelease(release string) ([]Book, error) {

	FoundBooks := make([]Book, 0, len(lib.Books))

	for _, book := range lib.Books {
		if book.ReleaseDate == release {
			FoundBooks = append(FoundBooks, book)
		}
	}
	if len(FoundBooks) == 0 {
		return nil, errors.New("Нет книг с таким годом выпуска ")
	}
	return FoundBooks, nil
}
