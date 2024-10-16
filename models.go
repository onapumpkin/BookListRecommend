package main

import "time"

type author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type userInfo struct {
	ID               string   `json:"id"`
	Username         string   `json:"username"`
	PreferredAuthors []string `json:"preferredAuthors"`
}

type book struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Author          author    `json:"author"`
	Description     string    `json:"description"`
	Genre           string    `json:"genre"`
	PublicationDate time.Time `json:"publicationDate"`
}

type userBookList struct {
	Username         string `json:"username"`
	RecommendedBooks []book `json:"recommendedBooks"`
}

type pairedReaderList struct {
	PairedUser1 *userBookList `json:"firstBookList"`
	PairedUser2 *userBookList `json:"secondBookList"`
}
