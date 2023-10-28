package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ondbyte/orderbook"
	"github.com/ondbyte/orderbook/example/example-web-app/static"
)

func main() {
	ob := orderbook.NewOrderBook()

	router := gin.Default()
	router.GET("/new-order", static.NewOrderPage())

	router.POST("/place-order")
}
