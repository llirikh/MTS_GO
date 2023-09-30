package pkg

import "errors"

type Storage interface {
	GetBook(uint64) (*Book, error)
	AddBook(uint64, *Book)
}

// MapStorage - data
type MapStorage struct {
	data map[uint64]*Book
}

func NewMapStorage() Storage {
	return &MapStorage{map[uint64]*Book{}}
}

func (mapStorage *MapStorage) GetBook(id uint64) (*Book, error) {
	if _, ok := mapStorage.data[id]; !ok {
		return &Book{}, errors.New("book is not found")
	}
	return mapStorage.data[id], nil
}

func (mapStorage *MapStorage) AddBook(id uint64, book *Book) {
	mapStorage.data[id] = book
}

// SliceStorage - data
type SliceStorage struct {
	data   []*Book
	dataId []uint64
}

func NewSliceStorage() Storage {
	return &SliceStorage{[]*Book{}, []uint64{}}
}

func (sliceStorage *SliceStorage) GetBook(id uint64) (*Book, error) {
	for i := 0; i < len(sliceStorage.data); i++ {
		if sliceStorage.dataId[i] == id {
			return sliceStorage.data[i], nil
		}
	}
	return &Book{}, errors.New("book is not found")
}

func (sliceStorage *SliceStorage) AddBook(id uint64, book *Book) {
	sliceStorage.data = append(sliceStorage.data, book)
	sliceStorage.dataId = append(sliceStorage.dataId, id)
}
