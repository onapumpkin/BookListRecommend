package main

import (
	"sort"
	"time"

	"golang.org/x/exp/rand"
)

func checkPublicationDate(bookPublicationDate int) bool {
	t := time.Now()
	currentYear := t.Year()
	return (currentYear - bookPublicationDate) <= 2
}

// RemoveOldPublications, removes books older than 2 year
func RemoveOldPublications(bookList []book) []book {
	freshestBookList := []book{}
	for _, book := range bookList {
		// checks to make sure book is published in last 2 years
		if checkPublicationDate(book.PublicationDate.Year()) {
			freshestBookList = append(freshestBookList, book)
		}
	}
	return freshestBookList
}

// GetListOfPreferredAuthors, returns a list of author structs using author name only
func GetListOfPreferredAuthors(authors []string) []author {
	verifiedAuthors := []author{}
	for i, authorName := range authors {
		// make sure to only include 5
		if i < 5 {
			potentialAuthor := FetchAuthorsByName(authorName)
			// there is only one author with that name add it
			if len(potentialAuthor) == 1 {
				verifiedAuthors = append(verifiedAuthors, potentialAuthor[0])
				// if there are multiples with same name, choose author with the most books
			} else if len(potentialAuthor) > 1 {
				var correctAuthor author
				for _, auth := range potentialAuthor {
					books := FetchBooksByAuthorID(auth.ID)
					tempLength := len(books)
					var amtBooks int
					if correctAuthor.ID == "" {
						correctAuthor = auth
						amtBooks = tempLength
					} else if tempLength > amtBooks {
						correctAuthor = auth
						amtBooks = tempLength
					}
				}
				verifiedAuthors = append(verifiedAuthors, correctAuthor)
			}
		}
	}
	return verifiedAuthors
}

// GetGenresMap, returns a list of genres given a list of authors
func GetGenresMap(authors []author) map[string]int {
	genres := make(map[string]int)
	for _, author := range authors {
		// for each author, grab the books and create the genre map per book genre
		books := FetchBooksByAuthorID(author.ID)
		for _, book := range books {
			if genres[book.Genre] == 0 {
				genres[book.Genre] = 1
			} else {
				genres[book.Genre] += 1
			}
		}
	}
	return SortGenreMap(genres)
}

// SortGenreMap, sorts genre map with most
func SortGenreMap(genres map[string]int) map[string]int {
	var indexes []string
	for i := range genres {
		indexes = append(indexes, i)
	}

	sort.SliceStable(indexes, func(i, j int) bool {
		return genres[indexes[i]] > genres[indexes[j]]
	})
	return genres
}

// FindCommonGenre compare two sorted maps of genres/bookcount and return top common genre
func FindCommonGenre(genres1, genres2 map[string]int) string {
	var topCommonGenre string
	for genre1 := range genres1 {
		for genre2 := range genres2 {
			if genre1 == genre2 {
				topCommonGenre = genre1
			}
		}
	}
	return topCommonGenre
}

// GetRecommendedBooks
func GetRecommendedBooks(genre string) []book {
	// get list of books by genre
	books := FetchBooksByGenre(genre)

	// only recommend books still in print
	newerBooks := RemoveOldPublications(books)

	// sort newest to oldest by publication date
	sort.Slice(newerBooks, func(i, j int) bool {
		return newerBooks[i].PublicationDate.After(newerBooks[j].PublicationDate)
	})

	randInt := rand.Intn(4) + 1
	return newerBooks[:randInt]
}
