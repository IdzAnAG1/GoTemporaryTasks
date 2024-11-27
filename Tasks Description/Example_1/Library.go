package Example_1

func NewLibrary() *Library {
	return &Library{
		Books:              map[string]*Book{},
		Clients:            map[string]*Client{},
		MaxBooksPerMembers: 5,
	}
}
