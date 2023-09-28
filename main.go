package main

import (
	"MTS_GO/pkg"
	"fmt"
)

func main() {
	// create array of Books
	book := []pkg.Book{
		pkg.NewBook("Герой нашего времени", "Михаил Лермонтов", "Лабиринт"),
		pkg.NewBook("Жёлтая стрела", "Виктор Пелевин", "Лабиринт"),
		pkg.NewBook("Путешествие в Икстлан", "Карлос Кастенеда", "София"),
		pkg.NewBook("Pale Fire", "Vladimir Nabokov", "G. P. Putnam's Sons"),
		pkg.NewBook("Сила и слава", "Грэм Грин", "Художественная литература"),
	}

	// test Library on MapStorage and PolynomialHash
	library1 := pkg.NewLibrary(pkg.NewMapStorage(), pkg.PolynomialHash())

	for i := 0; i < len(book); i++ {
		library1.AddBook(book[i])
	}

	fmt.Println()
	fmt.Println("Library on MapStorage and PolynomialHash:")
	// book would be found
	fmt.Println(library1.GetBook("Путешествие в Икстлан"))
	fmt.Println(library1.GetBook("Pale Fire"))
	// book wouldn't be found
	fmt.Println(library1.GetBook("Содом и Гоморра"))

	// test Library on SliceStorage and ReversePolynomialHash
	library2 := pkg.NewLibrary(pkg.NewSliceStorage(), pkg.ReversePolynomialHash())

	for i := 0; i < len(book); i++ {
		library2.AddBook(book[i])
	}

	fmt.Println()
	fmt.Println("Library on SliceStorage and ReversePolynomialHash:")
	// book would be found
	fmt.Println(library2.GetBook("Сила и слава"))
	fmt.Println(library2.GetBook("Жёлтая стрела"))
	// book wouldn't be found
	fmt.Println(library2.GetBook("Заводной апельсин"))
}
