package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SizeOptions struct {
	Sizes []int `json:"sizes"`
}

var sizeOptions = SizeOptions{
	Sizes: []int{9, 13, 19},
}

func GetSizeOptions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sizeOptions)
}
