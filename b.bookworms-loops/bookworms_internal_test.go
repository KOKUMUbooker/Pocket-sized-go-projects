package main

import "testing"

var oryxAndCrake = Book {
    Author: "Margaret Atwood",
    Title: "Oryx and Crake",
};
	
var	handmaidsTale = Book {
	Author: "Margaret Atwood",
    Title: "The Handmaid's Tale",
}

var janeEyre = Book {
    Author: "Charlotte Brontë",
    Title: "Jane Eyre",
}

var theBellJar =Book {
    Author: "Sylvia Plath",
    Title: "The Bell Jar",
}

// ============================================= Test #1 =======================================================
type testCase struct {
	bookwormsFile string
	want []Bookworm
	wantErr bool
}

var tests = map[string]testCase {
	"file exists": {
		bookwormsFile: "testdata/bookworms.json",
		want: []Bookworm{
			{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
			{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
		},
		wantErr: false,
		},	
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want: nil,
		wantErr: true,
	},
	"invalid JSON": {
		bookwormsFile: "testdata/invalid.json",
		want:nil,
		wantErr:true,
	},
};

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
	return false 
	}

	for i := range books {
		if books[i] != target[i] { 
			return false
		}
	}

	return true 
}

// equalBookworms is a helper to test the equality of two lists of Bookworms
func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms) != len(target) {
		return false 
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name { 
			return false
		}
		
		if !equalBooks(bookworms[i].Books, target[i].Books) { 
			return false
		}
	}
	
	return true 
}

func TestLoadBookworms(t *testing.T) {
	for name,testCase := range tests {
		t.Run(name,func(t *testing.T) {
			got,err := LoadBookworms(testCase.bookwormsFile);
			
			if err != nil && !testCase.wantErr { 
				t.Fatalf("expected an error %s, got one", err.Error())
			}

			if err == nil && testCase.wantErr { 
				t.Fatalf("expected an error, got none")
			}

			if !equalBookworms(got, testCase.want) { 
				t.Fatalf("different result: got %v, expected %v", got, testCase)
			}

		})
	}
}

// ============================================= Test #2 =======================================================
// equalBooksCount is a helper to test the equality of two maps of books cou
func equalBooksCount(got, want map[Book]uint) bool {
	if len(got) != len(want) { 
		return false
	}
	
	for book, targetCount := range want { 
		count, ok := got[book] 
		if !ok || targetCount != count { 
			return false 
		}
	}

	return true 
}

type bcTestCase struct {
	input []Bookworm;
	want map[Book]uint;
}

func TestBooksCount(t *testing.T) {
	tt := map[string] bcTestCase{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}}, 
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, janeEyre:1,oryxAndCrake:1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want: map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm {
				{Name:"John",Books:[]Book{}},
				{Name:"Jane",Books:[]Book{}},
			},
			want: map[Book]uint{},
		},
		"bookworm with twice the same book": {
			input: []Bookworm {
				{Name:"Lobster",Books:[]Book{handmaidsTale,handmaidsTale}},
			},
			want: map[Book]uint{handmaidsTale:2},
		},
	};

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(tc.want, got) { 
				t.Fatalf("got a different list of books: %v, expected %v", got ,tc.want)
			}
		})
	}
}

// ============================================= Test #3 =======================================================
type fcbTestCase struct {
	bookworms []Bookworm
	want []Book
}

func equalComBooks(got,want []Book) bool{
	if len(got) != len(want){
		return false
	}

	for i,book := range got {
		if want[i] != book {
			return false
		}
	}

	return  true;
}

func TestFindCommonBooks(t *testing.T) {
	var testCases map[string]fcbTestCase = map[string]fcbTestCase{
		"Everyone has read the same books": {
			bookworms: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}}, 
				{Name: "Difa", Books: []Book{handmaidsTale, theBellJar}}, 
			},
			want: []Book{handmaidsTale,theBellJar},
		},
		"People have completely different lists": {
			bookworms: []Bookworm{
				{Name:"John",Books:[]Book{handmaidsTale}},
				{Name:"Jane",Books:[]Book{theBellJar}},
			},
			want: []Book{},
		},
		"More than 2 bookworms have a book in common": {
			bookworms: []Bookworm{
				{Name:"Ralph",Books:[]Book{handmaidsTale,theBellJar,janeEyre,oryxAndCrake}},
				{Name:"Sue",Books:[]Book{handmaidsTale,theBellJar,janeEyre}},
			},
			want: []Book{handmaidsTale,theBellJar,janeEyre},
		},
		"One bookworm has no books (oh the sadness!)": {
			bookworms: []Bookworm{
				{Name:"Don",Books:[]Book{}},
				{Name:"Polo",Books:[]Book{theBellJar}},
			},
			want: []Book{},
		},
		"Nobody has any book (oh the agony!)": {
			bookworms: []Bookworm{
				{Name:"Carly",Books:[]Book{}},
				{Name:"Forlan",Books:[]Book{}},
			},
			want: []Book{},
		},
	}; 
	
	for name,testCase := range testCases {
		t.Run(name,func(t *testing.T) {
			got := findCommonBooks(testCase.bookworms);

			if !equalComBooks(got,testCase.want) {
				t.Fatalf("Expected : %v, got : %v ",testCase.want,got);
			}
		})
	}
}