package Controllers

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/the-r3aper7/stock-price-server/Helpers"
)

func GetMultiplePrice(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	tickers := strings.Split(c.Param("tickers"), ",")
	size := len(tickers)
	chj := make(chan Helpers.ChannelJsonData, size)
	results := []Helpers.ChannelJsonData{}

	for _, ticker := range tickers {
		ticker := strings.Trim(ticker, " ")
		go Helpers.MakeRequest(ticker, chj)
	}

	for range tickers {
		results = append(results, <-chj)
	}
	// sorting so tha after completeion of go routines it gives always the same results
	sort.Slice(results, func(i, j int) bool {
		return results[i].Data.Symbol < results[j].Data.Symbol
	})

	c.IndentedJSON(http.StatusOK, results)
}
