package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ondbyte/orderbook/example/example-web-app/orders"
	"github.com/ondbyte/orderbook/example/example-web-app/static"
)

func Run(addr string) error {
	router := gin.Default()
	handler := orders.NewOrderHandler(orders.NewService())
	orderRoute := "/order"
	router.GET(orderRoute, static.NewOrderPage(orderRoute))
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, orderRoute)
	})
	router.POST(orderRoute, handler.HandleOrder)
	router.GET("/updates/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if ctx.GetHeader("Upgrade") == "websocket" {
			handler.HandleOrderUpdates(ctx)
		} else {
			static.NewOrderUpdatesPage("/updates/" + id)(ctx)
		}

	})
	return router.Run(addr)
}
