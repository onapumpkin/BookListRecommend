package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	bestieUsername1 = "onapumpkin"
	bestieUsername2 = "notonapumpkin"
	genres          = []string{"fantasy", "horror", "mystery", "cop drama"}
	listOfAuthors   map[string]string
)

func TestFetchSharedBookList(t *testing.T) {
	testEng := MockEngine(t)
	w := httptest.NewRecorder()
	testSetup()
	request, err := http.NewRequest(http.MethodGet, "/paired_recommend/12345/56789", nil)
	assert.NoError(t, err)
	testEng.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Printf("Response: %+v\n", w.Body.String())

	var response pairedReaderList
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, response.PairedUser1.Username, bestieUsername1)
	assert.Equal(t, response.PairedUser2.Username, bestieUsername2)
	assert.Greater(t, len(response.PairedUser1.RecommendedBooks), 1)
	assert.Greater(t, len(response.PairedUser2.RecommendedBooks), 1)
}

func MockEngine(t *testing.T) *gin.Engine {
	router := gin.Default()
	router.GET("/paired_recommend/:id1/:id2", FetchSharedBookList)
	return router
}

func testSetup() {
	MapOfAuthorToID()
	mockFetchUserInfoByUserID()
	mockFetchBooksByGenre()
	mockFetchBooksByAuthorID()
	mockFetchAuthorsByName()
}

func MapOfAuthorToID() {
	listOfAuthors = make(map[string]string)
	listOfAuthors["1"] = "Robin McKinley"
	listOfAuthors["2"] = "Robin McKinley"
	listOfAuthors["3"] = "Test Author1"
	listOfAuthors["4"] = "Upset Cow"
	listOfAuthors["5"] = "Random Name"
	listOfAuthors["6"] = "S Bottoms"
	listOfAuthors["7"] = "S Bottoms"
	listOfAuthors["8"] = "TS Lowe"
	listOfAuthors["9"] = "Help Text"
	listOfAuthors["10"] = "Onemore Author"
	listOfAuthors["11"] = "Extra Author"
}

func mockFetchUserInfoByUserID() {
	FetchUserInfoByUserID = func(id string) (userInfo, error) {
		if id == "12345" {
			return userInfo{ID: id, Username: bestieUsername1, PreferredAuthors: []string{"Robin McKinley", "Test Author1", "Upset Cow", "Random Name", "S Bottoms"}}, nil
		} else if id == "56789" {
			return userInfo{ID: id, Username: bestieUsername2, PreferredAuthors: []string{"Test Author1", "Random Name", "TS Lowe", "Help Text", "Onemore Author", "Extra Author"}}, nil
		}
		return userInfo{}, fmt.Errorf("No user information was found")
	}
}

func mockFetchAuthorsByName() {
	FetchAuthorsByName = func(name string) []author {
		authors := []author{}
		for i, authorName := range listOfAuthors {
			if authorName == name {
				authors = append(authors, author{ID: i, Name: authorName})
			}
		}
		return authors
	}
}

func mockFetchBooksByGenre() {
	FetchBooksByGenre = func(genre string) []book {
		books := []book{}
		if strings.ToLower(genre) == "fantasy" {
			for i := 0; i < 50; i++ {
				id := uuid.New()
				// randomize which author without equal ammounts
				index := strconv.Itoa(rand.IntN(100)%11 + 1)
				a := listOfAuthors[index]
				randomYear := 2020 + (rand.IntN(10) % 5)
				book := book{
					ID:    id.String(),
					Title: fmt.Sprintf("Book of Frufru %v", i),
					Author: author{
						ID:   index,
						Name: a,
					},
					Description:     "",
					PublicationDate: time.Date(randomYear, 12, i%28+1, 0, 0, 0, 0, time.UTC),
					Genre:           "fantasy",
				}
				books = append(books, book)
			}
		} else if strings.ToLower(genre) == "horror" {
			for i := 0; i < 50; i++ {
				id := uuid.New()
				// randomize which author without equal ammounts
				index := strconv.Itoa(rand.IntN(100)%11 + 1)
				a := listOfAuthors[index]
				randomYear := 2020 + (rand.IntN(10) % 5)
				book := book{
					ID:    id.String(),
					Title: fmt.Sprintf("Book of Horror %v", i),
					Author: author{
						ID:   index,
						Name: a,
					},
					Description:     "",
					PublicationDate: time.Date(randomYear, 12, i%28+1, 0, 0, 0, 0, time.UTC),
					Genre:           "horror",
				}
				books = append(books, book)
			}
		} else if strings.ToLower(genre) == "mystery" {
			for i := 0; i < 50; i++ {
				id := uuid.New()
				// randomize which author without equal ammounts
				index := strconv.Itoa(rand.IntN(100)%11 + 1)
				a := listOfAuthors[index]
				randomYear := 2020 + (rand.IntN(10) % 5)
				book := book{
					ID:    id.String(),
					Title: fmt.Sprintf("Mysteries of Doom %v", i),
					Author: author{
						ID:   index,
						Name: a,
					},
					Description:     "",
					PublicationDate: time.Date(randomYear, 12, i%28+1, 0, 0, 0, 0, time.UTC),
					Genre:           "mystery",
				}
				books = append(books, book)
			}
		} else if strings.ToLower(genre) == "cop drama" {
			for i := 0; i < 50; i++ {
				id := uuid.New()
				// randomize which author without equal ammounts
				index := strconv.Itoa(rand.IntN(100)%11 + 1)
				a := listOfAuthors[index]
				randomYear := 2020 + (rand.IntN(10) % 5)
				book := book{
					ID:    id.String(),
					Title: fmt.Sprintf("Cops of Time Square %v", i),
					Author: author{
						ID:   index,
						Name: a,
					},
					Description:     "",
					PublicationDate: time.Date(randomYear, 12, i%28+1, 0, 0, 0, 0, time.UTC),
					Genre:           "cope drama",
				}
				books = append(books, book)
			}
		}
		return books
	}
}

func mockFetchBooksByAuthorID() {
	FetchBooksByAuthorID = func(authorID string) []book {
		books := []book{}
		for i := 1; i < 12; i++ {
			index := strconv.Itoa(i)
			for j := 0; j < 100; j++ {
				id := uuid.New()
				randomYear := 2020 + (rand.IntN(10) % 5)
				book := book{
					ID:    id.String(),
					Title: fmt.Sprintf("Book of Frufru %v", i),
					Author: author{
						ID:   index,
						Name: listOfAuthors[index],
					},
					Description:     "",
					PublicationDate: time.Date(randomYear, 12, i%28+1, 0, 0, 0, 0, time.UTC),
					Genre:           genres[i%4],
				}
				books = append(books, book)
			}
		}
		return books
	}
}
