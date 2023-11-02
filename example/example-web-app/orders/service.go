package orders

import (
	"github.com/ondbyte/orderbook"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Service struct {
	// order book
	ob *orderbook.OrderBook
}

func NewService() *Service {
	return &Service{
		ob: orderbook.NewOrderBook(),
	}
}

func (s *Service) PlaceMarketOrder(side orderbook.Side, qty decimal.Decimal) (*MarketOrderResponse, error) {
	bought, partiallybought, partialQty, left, err := s.ob.ProcessMarketOrder(side, qty)
	if err != nil {
		return nil, err
	}
	txns := make([]*ProcessedTransaction, 0)
	for _, order := range bought {
		txns = append(txns, &ProcessedTransaction{
			Id:       order.ID(),
			Quantity: order.Quantity(),
			Price:    order.Price(),
		})
	}
	return &MarketOrderResponse{
		Left: left,
		OrderResponse: &OrderResponse{
			Side:                  side,
			ProcessedQuantity:     partialQty,
			ProcessedTransactions: txns,
			PartiallyProcessedTransaction: &ProcessedTransaction{
				Id:       partiallybought.ID(),
				Quantity: partiallybought.Quantity(),
				Price:    partiallybought.Price(),
			},
		},
	}, nil
}

func (s *Service) PlaceLimitOrder(side orderbook.Side, qty decimal.Decimal, price decimal.Decimal) (*LimitOrderResponse, error) {
	id := uuid.NewV4().String()
	bought, partiallybought, partialQty, err := s.ob.ProcessLimitOrder(side, id, qty, price)
	if err != nil {
		return nil, err
	}
	txns := make([]*ProcessedTransaction, 0)
	for _, order := range bought {
		txns = append(txns, &ProcessedTransaction{
			Id:       order.ID(),
			Quantity: order.Quantity(),
			Price:    order.Price(),
		})

		listener, ok := orderListeners[order.ID()]
		if ok {
			listener.WriteJSON(order)
		}
	}
	var partialTxn *ProcessedTransaction
	if partiallybought != nil {
		partialTxn = &ProcessedTransaction{
			Id:       partiallybought.ID(),
			Quantity: partiallybought.Quantity(),
			Price:    partiallybought.Price(),
		}
	}
	return &LimitOrderResponse{
		Id: id,
		OrderResponse: &OrderResponse{
			Side:                          side,
			RequestedQuantity:             qty,
			ProcessedQuantity:             partialQty,
			ProcessedTransactions:         txns,
			PartiallyProcessedTransaction: partialTxn,
		},
	}, nil
}
