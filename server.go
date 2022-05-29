package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-r3aper7/stock-price-server/Controllers"
)

func main() {
	r := gin.Default()
	r.GET("/:ticker", Controllers.GetPrice)
	r.GET("/m/:tickers", Controllers.GetMultiplePrice)
	r.GET("/favicon.ico", func(ctx *gin.Context) { ctx.File("./favicon.ico") })
	r.Run()
}
