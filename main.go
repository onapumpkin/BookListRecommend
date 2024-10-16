package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	setupDB()
	bookEng := gin.Default()
	bookEng.GET("/paired_recommend/:id1/:id2", FetchSharedBookList)
	bookEng.Run()
}

func FetchSharedBookList(c *gin.Context) {
	userID1 := c.Param("id1")
	userID2 := c.Param("id2")

	userInfo1, err := FetchUserInfoByUserID(userID1)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	userInfo2, err := FetchUserInfoByUserID(userID2)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user1FavAuthors := GetListOfPreferredAuthors(userInfo1.PreferredAuthors)
	user1Genres := GetGenresMap(user1FavAuthors)

	user2FavAuthors := GetListOfPreferredAuthors(userInfo2.PreferredAuthors)
	user2Genres := GetGenresMap(user2FavAuthors)

	commonGenre := FindCommonGenre(user1Genres, user2Genres)

	pairedList := pairedReaderList{
		PairedUser1: &userBookList{
			Username:         userInfo1.Username,
			RecommendedBooks: GetRecommendedBooks(commonGenre),
		},
		PairedUser2: &userBookList{
			Username:         userInfo2.Username,
			RecommendedBooks: GetRecommendedBooks(commonGenre),
		},
	}

	c.JSON(http.StatusOK, pairedList)
}
