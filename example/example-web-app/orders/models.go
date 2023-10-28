package orders

import (
	"github.com/ondbyte/orderbook"
	"github.com/shopspring/decimal"
)

type OrderDetails struct {
	Id       string          `json:"id" form:"id" binding:"required"`
	Side     orderbook.Side  `json:"side" form:"side" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" form:"quantity" binding:"required"`
	Price    decimal.Decimal `json:"price" form:"price" binding:"required"`
}

type TransactionsDetails struct {
	Side orderbook.Side `json:"side" form:"side" binding:"required"`
	// total qty we bought/sold
	Quantity decimal.Decimal `json:"quantity" form:"quantity" binding:"required"`
	// number of stocks we are unable to buy
	Left decimal.Decimal `json:"left" form:"left" binding:"required"`
	// complete orders we bought/sold
	Transactions []*TransactionDetails `json:"transactions" form:"transactions" binding:"required"`
	// partial orders bought or sold
	PartialTransaction *TransactionDetails `json:"partialTransaction" form:"partialTransaction" binding:"required"`
}

type TransactionDetails struct {
	Id       string          `json:"id" form:"id" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" form:"quantity" binding:"required"`
	Price    decimal.Decimal `json:"price" form:"price" binding:"required"`
}
