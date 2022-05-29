package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/the-r3aper7/stock-price-server/Helpers"
)

func GetPrice(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	chj := make(chan Helpers.ChannelJsonData)
	ticker := c.Param("ticker")

	go Helpers.MakeRequest(ticker, chj)

	c.IndentedJSON(http.StatusOK, <-chj)
}
