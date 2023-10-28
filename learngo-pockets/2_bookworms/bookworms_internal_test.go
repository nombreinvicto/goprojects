package main

import (
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood",
		Title: "The Handmaid's Tale"}
	oryxAndCrake = Book{Author: "Margaret Atwood",
		Title: "Oryx and Crake"}
	theBellJar = Book{Author: "Sylvia Plath",
		Title: "The Bell Jar"}
	janeEyre = Book{Author: "Charlotte Brontë",
		Title: "Jane Eyre"}
)

func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		// seems like it is doing a deepequal comparison
		// because each Book is a struct, so we'r doing struct
		//comparison over here
		if books[i] != target[i] {
			return false
		}
	}
	return true
}

func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		// now check if all Books of a bookwork match
		// that of the target
		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}
func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata\\bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesnt exist": {
			bookwormsFile: "testdata\\no_file.json",
			want:          nil,
			wantErr:       true,
		},
	}

	// range over all scenarios
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)

			// test gave an error, but was NOT expecting one
			// e.g case when I have valid filepath but still got error
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())

			}

			// test didnt give an error but was expecting one
			// e.g case when i have invalid filepath, so expecting error but didnt get one
			if err == nil && testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}

			// will reach here if no probs in above steps
			// so now check what i got is valid
			// if it was invalid filepath, got is nil by now
			// got is either []Bookwork or nil
			if !equalBookworms(got, testCase.want) {
				t.Fatalf("different results: got %v, expected %v", got, testCase.want)
			}

		})
	}
}

// equalBooksCount is a helper that tests whether
// 2 booksCount maps are equal or not
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

func TestBooksCount(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  map[Book]uint
	}

	tests := map[string]testCase{
		"nominal_use_case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no_bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(got, tc.want) {
				t.Fatalf("got different list of books, got: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			if !equalBooks(tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}
