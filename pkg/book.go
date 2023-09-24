package pkg

type Book struct {
	title     string
	author    string
	publisher string
}

func NewBook(title, author, publisher string) *Book {
	return &Book{title, author, publisher}
}

func (book *Book) GetTitle() *string {
	return &book.title
}

func (book *Book) GetAuthor() *string {
	return &book.author
}

func (book *Book) GetPublisher() *string {
	return &book.publisher
}
