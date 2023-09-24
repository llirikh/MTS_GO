package main

import (
	"MTS_GO/pkg"
	"fmt"
)

func main() {
	// create array of Books
	book := []*pkg.Book{
		pkg.NewBook("Герой нашего времени", "Михаил Лермонтов", "Лабиринт"),
		pkg.NewBook("Жёлтая стрела", "Виктор Пелевин", "Лабиринт"),
		pkg.NewBook("Путешествие в Икстлан", "Карлос Кастенеда", "София"),
		pkg.NewBook("Pale Fire", "Vladimir Nabokov", "G. P. Putnam's Sons"),
		pkg.NewBook("Сила и слава", "Грэм Грин", "Художественная литература"),
	}

	// test Library on MapStorage and PolynomialHash
	storage1 := pkg.NewMapStorage()
	hash1 := pkg.PolynomialHash()

	library1 := pkg.NewLibrary(storage1, hash1)

	for i := 0; i < len(book); i++ {
		library1.AddBook(*book[i])
	}

	fmt.Println()
	fmt.Println("Library on MapStorage and PolynomialHash:")
	fmt.Println(library1.GetBook("Путешествие в Икстлан"))
	fmt.Println(library1.GetBook("Pale Fire"))

	// test Library on SliceStorage and ReversePolynomialHash
	storage2 := pkg.NewSliceStorage()
	hash2 := pkg.ReversePolynomialHash()

	library2 := pkg.NewLibrary(storage2, hash2)

	for i := 0; i < len(book); i++ {
		library2.AddBook(*book[i])
	}

	fmt.Println()
	fmt.Println("Library on SliceStorage and ReversePolynomialHash:")
	fmt.Println(library2.GetBook("Сила и слава"))
	fmt.Println(library2.GetBook("Жёлтая стрела"))
}
