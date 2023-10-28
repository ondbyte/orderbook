package orders

import (
	"github.com/ondbyte/orderbook"
	"github.com/shopspring/decimal"
)

type Service struct {
	// order book
	ob *orderbook.OrderBook
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) PlaceOrder(order *OrderDetails) (*TransactionsDetails, error) {
	// if the price is empty its market order
	if order.Price.Equal(decimal.Decimal{}) {
		return s.PlaceMarketOrder(order.Side, order.Quantity)
	}
	return s.
}

func (s *Service) PlaceMarketOrder(side orderbook.Side, qty decimal.Decimal) (*TransactionsDetails, error) {
	bought, partiallybought, partialQty, left, err := s.ob.ProcessMarketOrder(side, qty)
	if err != nil {
		return nil, err
	}
	txns := make([]*TransactionDetails, 0)
	for _, order := range bought {
		txns = append(txns, &TransactionDetails{
			Id:       order.ID(),
			Quantity: order.Quantity(),
			Price:    order.Price(),
		})
	}
	return &TransactionsDetails{
		Side:         side,
		Quantity:     partialQty,
		Left:         left,
		Transactions: txns,
		PartialTransaction: &TransactionDetails{
			Id:       partiallybought.ID(),
			Quantity: partiallybought.Quantity(),
			Price:    partiallybought.Price(),
		},
	}, nil
}

func (s *Service) PlaceMarketOrder(side orderbook.Side, qty decimal.Decimal) (*TransactionsDetails, error) {
	bought, partiallybought, partialQty, left, err := s.ob.ProcessMarketOrder(side, qty)
	if err != nil {
		return nil, err
	}
	txns := make([]*TransactionDetails, 0)
	for _, order := range bought {
		txns = append(txns, &TransactionDetails{
			Id:       order.ID(),
			Quantity: order.Quantity(),
			Price:    order.Price(),
		})
	}
	return &TransactionsDetails{
		Side:         side,
		Quantity:     partialQty,
		Left:         left,
		Transactions: txns,
		PartialTransaction: &TransactionDetails{
			Id:       partiallybought.ID(),
			Quantity: partiallybought.Quantity(),
			Price:    partiallybought.Price(),
		},
	}, nil
}
