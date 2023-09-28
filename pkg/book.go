package pkg

type Book struct {
	Title     string
	author    string
	publisher string
}

func NewBook(title, author, publisher string) Book {
	return Book{title, author, publisher}
}

func (book Book) GetTitle() string {
	return book.Title
}

func (book Book) GetAuthor() string {
	return book.author
}

func (book Book) GetPublisher() string {
	return book.publisher
}
