package pkg

type Library struct {
	storage Storage
	getId   IdGenerator
}

func NewLibrary(storage Storage, getId IdGenerator) Library {
	return Library{storage, getId}
}

func (lib *Library) GetBook(title string) Book {
	id := lib.getId(title)
	return lib.storage.GetBook(id)
}

func (lib *Library) AddBook(book Book) {
	id := lib.getId(book.title)
	lib.storage.AddBook(id, book)
}
