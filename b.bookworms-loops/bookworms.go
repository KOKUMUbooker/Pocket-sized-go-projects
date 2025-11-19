package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Book struct {
	Author string `json:"author"`; 
	Title string `json:"title"`;
}

type Bookworm struct {
	Name string `json:"name"`;
	Books []Book `json:"books"`;
}

// Reads the file and returns the list of bookworms
func LoadBookworms(filePath string) ([]Bookworm,error) {
	f,err := os.Open(filePath);
	if err != nil {
		return nil,err
	}

	defer f.Close();
	// Initialize the file in which the file will be decoded into
	var bookworms []Bookworm;

	// Decode file contents & store them into the variable bookworms
	err = json.NewDecoder(f).Decode(&bookworms);
	if err != nil {
		return  nil,err;
	}

	return  bookworms,nil;
}


// booksCount registers all the books and their occurrences from the bookwor
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint) 
	for _, bookworm := range bookworms { 
		for _, book := range bookworm.Books {
			count[book]++ 
		}
	}
	return count
}

// findCommonBooks returns books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	var commonBooks []Book;
	booksOnShelves := booksCount(bookworms);

	for book, count := range booksOnShelves { 
		if count > 1 { 
			commonBooks = append(commonBooks, book)
		}
	}

	return commonBooks;
}

// sortBooks sorts the books by Author and then Title.
// sort.Slice(mySlice, func(i,j int)bool) will modify mySlice
func SortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool { 
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author 
		}
		return books[i].Title < books[j].Title 
	});

	return books
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}