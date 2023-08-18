package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type gameInfo struct {
	ID             string  `json:"id"`
	BlackPositions [][]int `json:"blackPositions"`
	WhitePositions [][]int `json:"whitePositions"`
	Dimensions     int     `json:"dimensions"`
	WhiteToPlay    bool    `json:"whiteToPlay"`
}

var gameInfos = []gameInfo{
	{ID: "1", BlackPositions: [][]int{{2, 1}}, WhitePositions: [][]int{{1, 2}, {1, 1}}, Dimensions: 9, WhiteToPlay: false},
	{ID: "1", BlackPositions: [][]int{{5, 5}}, WhitePositions: [][]int{{4, 4}}, Dimensions: 13, WhiteToPlay: true},
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)

	router.Run("localhost:8080")
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gameInfos)
}
