package Example_2

type Book struct {
	Title       string `json:"title"`        // Название книги
	Author      string `json:"author"`       // Автор книги
	ReleaseDate int    `json:"release_date"` // Год выпуска книги
	Status      string `json:"Status"`       // Статус в котором находится книга (Выдана или Не выдана)
}
type Library struct {
	Books map[string]Book `json:"books"` // Ключ для карты Название книги
}
