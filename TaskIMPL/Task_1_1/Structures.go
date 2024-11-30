package Task_1_1

type Book struct {
	Title       string
	Author      string
	ReleaseDate string
	Status      bool
}
type Library struct {
	Books map[string]Book // Ключ для карты Название книги
}
