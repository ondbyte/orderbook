package orders

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ondbyte/orderbook"
	"github.com/shopspring/decimal"
)

type OrderRequest struct {
	Side     orderbook.Side  `json:"side" form:"side" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" form:"quantity" binding:"required"`
	Price    decimal.Decimal `json:"price" form:"price" binding:"required"`
}

func NewOrderRequestFromReq(ctx *gin.Context) (*OrderRequest, error) {
	m := map[string]string{
		"side":     ctx.PostForm("side"),
		"quantity": ctx.PostForm("quantity"),
		"price":    ctx.PostForm("price"),
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	or := new(OrderRequest)
	err = json.Unmarshal(b, or)
	if err != nil {
		return nil, err
	}
	return or, nil
}

type OrderResponse struct {
	Side              orderbook.Side  `json:"side" form:"side" binding:"required"`
	RequestedQuantity decimal.Decimal `json:"requestedQuantity" form:"requestedQuantity" binding:"required"`
	// total qty we bought/sold
	ProcessedQuantity decimal.Decimal `json:"processedQuantity" form:"processedQuantity" binding:"required"`
	// complete orders we bought/sold
	ProcessedTransactions []*ProcessedTransaction `json:"processedTransactions" form:"processedTransactions" binding:"required"`
	// partial orders bought or sold
	PartiallyProcessedTransaction *ProcessedTransaction `json:"partiallyProcessedTransaction" form:"partiallyProcessedTransaction" binding:"required"`
}

type MarketOrderResponse struct {
	*OrderResponse
	// number of stocks we are unable to buy
	Left decimal.Decimal `json:"left" form:"left" binding:"required"`
}

type LimitOrderResponse struct {
	*OrderResponse
	// only limit order returns a id
	Id string `json:"id" form:"id" binding:"required"`
}

type ProcessedTransaction struct {
	Id       string          `json:"id" form:"id" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" form:"quantity" binding:"required"`
	Price    decimal.Decimal `json:"price" form:"price" binding:"required"`
}
