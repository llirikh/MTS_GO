package pkg

type Storage interface {
	GetBook(uint64) Book
	AddBook(uint64, Book)
}

// MapStorage - storage
type MapStorage struct {
	storage map[uint64]*Book
}

func NewMapStorage() *MapStorage {
	mapStorage := MapStorage{}
	mapStorage.storage = map[uint64]*Book{}
	return &mapStorage
}

func (mapStorage *MapStorage) GetBook(id uint64) Book {
	return *mapStorage.storage[id]
}

func (mapStorage *MapStorage) AddBook(id uint64, book Book) {
	mapStorage.storage[id] = &book
}

// SliceStorage - storage
type SliceStorage struct {
	storage []*Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{}
}

func (sliceStorage *SliceStorage) GetBook(id uint64) Book {
	return *sliceStorage.storage[id]
}

func (sliceStorage *SliceStorage) AddBook(id uint64, book Book) {
	if uint64(cap(sliceStorage.storage)) <= id {
		tempStorage := make([]*Book, id+1)
		copy(tempStorage, sliceStorage.storage)
		sliceStorage.storage = tempStorage
	}
	sliceStorage.storage[id] = &book
}
