package pkg

type Library struct {
	storage Storage
	getId   IdGenerator
}

func NewLibrary(storage Storage, getId IdGenerator) Library {
	return Library{storage, getId}
}

func (lib *Library) GetBook(title string) (Book, error) {
	id := lib.getId(&title)
	book, err := lib.storage.GetBook(id)
	return *book, err
}

func (lib *Library) AddBook(book Book) {
	id := lib.getId(&book.Title)
	lib.storage.AddBook(id, &book)
}
