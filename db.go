package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
)

var localDB *sql.DB

func setupDB() {
	// generic connection string
	connectionStr := "postgresql://postgres:password123@127.0.0.1/todos?sslmode=disable"

	localDB, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Local DB is setup, %v\n", localDB.Stats())
}

// FetchUserInfoByUserID, return user info based on userID
var FetchUserInfoByUserID = func(id string) (userInfo, error) {
	user := userInfo{}
	statement := `SELECT * FROM userinfo where userinfo.id=$1`
	err := localDB.QueryRow(statement, id).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// FetchBooksByGenre, return up to 50 books of a certain genre
var FetchBooksByGenre = func(genre string) []book {
	books := make([]book, 50)
	randomOffset := rand.IntN(100)
	statement := fmt.Sprintf(`SELECT * FROM books WHERE genre=$1 LIMIT 50 OFFSET %v`, randomOffset)
	rows, err := localDB.Query(statement, genre)
	if err != nil {
		return books
	}

	for rows.Next() {
		var b book
		err = rows.Scan(&b.ID, &b.Title, &b.Description, &b.Genre, &b.Author, &b.PublicationDate)
		if err != nil {
			return books
		}
		books = append(books, b)

	}
	if err = rows.Err(); err != nil {
		return books
	}

	return books
}

// FetchBooksByAuthorID, return books based on authorID
var FetchBooksByAuthorID = func(authorID string) []book {
	books := []book{}
	statement := `SELECT * FROM books WHERE author=$1`
	rows, err := localDB.Query(statement, authorID)
	if err != nil {
		return books
	}

	for rows.Next() {
		var b book
		err = rows.Scan(&b.ID, &b.Title, &b.Description, &b.Genre, &b.Author, &b.PublicationDate)
		if err != nil {
			return books
		}
		books = append(books, b)

	}

	return books
}

// FetchAuthorsByName, return authors based on author name
var FetchAuthorsByName = func(name string) []author {
	authors := []author{}
	statement := `SELECT id, name FROM authors WHERE name=$1`
	rows, err := localDB.Query(statement, name)
	if err != nil {
		return authors
	}
	for rows.Next() {
		var author author
		err = rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return authors
		}
		authors = append(authors, author)
	}

	return authors
}
