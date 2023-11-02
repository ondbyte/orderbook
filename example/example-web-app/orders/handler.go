package orders

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

type Handler struct {
	OrderBookService *Service
}

func NewOrderHandler(s *Service) *Handler {
	return &Handler{OrderBookService: s}
}

func (h *Handler) HandleOrder(ctx *gin.Context) {
	form, err := NewOrderRequestFromReq(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusTeapot, map[string]string{"error": err.Error()})
		return
	}
	if form.Price.Equal(decimal.Decimal{}) {
		details, err := h.OrderBookService.PlaceMarketOrder(form.Side, form.Quantity)
		if err != nil {
			log.Println(err)
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, details)
	} else {
		details, err := h.OrderBookService.PlaceLimitOrder(form.Side, form.Quantity, form.Price)
		if err != nil {
			log.Println(err)
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, details)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var orderListeners = map[string]*websocket.Conn{}

func (h *Handler) HandleOrderUpdates(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.Status(http.StatusTeapot)
		ctx.Error(fmt.Errorf("id param is required"))
		return
	}
	order := h.OrderBookService.ob.Order(id)
	if order == nil {
		ctx.Status(http.StatusNotFound)
		ctx.Error(fmt.Errorf("no order with id: %v", id))
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	conn.WriteJSON(order)
	orderListeners[id] = conn
}
