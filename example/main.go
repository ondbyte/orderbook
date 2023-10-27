package main

import (
	"log"

	"github.com/ondbyte/orderbook"
	"github.com/shopspring/decimal"
)

func main() {
	ob := orderbook.NewOrderBook()
	logger := log.Default()

	// person "A" sells his first order of 100 stock at a limit price of 100 per stock
	totalStocks := decimal.New(100, 0)
	pricePerStock := decimal.New(100, 0)
	buyOrders, partialBuyOrder, partialyQty, err := ob.ProcessLimitOrder(orderbook.Sell, "A-1", totalStocks, pricePerStock)
	if err != nil {
		logger.Panic(err)
	}
	// here we are using `buyOrders`, `partialBuyOrder` as the name of these variable is because when we sell the order buy orders which matched
	// for "A"s sell order will be returned
	// the above call to ProcessLimitOrder returns nil/zero values for all return values because there is no buy orders at market price or
	// limit price matching "A"s sell order
	if buyOrders != nil || len(buyOrders) > 0 {
		logger.Panicf("buyOrders for A order should be non nil")
	}
	if partialBuyOrder != nil {
		logger.Panicf("partialBuyOrder for A order should be nil")
	}
	if !partialyQty.Equal(decimal.Decimal{}) {
		logger.Panicf("partialyQty for A order should be zero but its %v", partialyQty)
	}

	// person "B" sells another 50 stock at a limit price of 99 per stock
	totalStocks = decimal.New(50, 0)
	pricePerStock = decimal.New(99, 0)
	buyOrders, partialBuyOrder, partialyQty, err = ob.ProcessLimitOrder(orderbook.Sell, "B-1", totalStocks, pricePerStock)
	// same checks as before
	if buyOrders != nil || len(buyOrders) > 0 {
		logger.Panicf("buyOrders for B order should be nil")
	}
	if partialBuyOrder != nil {
		logger.Panicf("partialBuyOrder for B order should be nil")
	}
	if !partialyQty.Equal(decimal.Decimal{}) {
		logger.Panicf("partialBuyOrder for B order should be nil")
	}
	// so as of now we have A selling 100 at 100
	// and B selling 50 at 99

	// now X buys 50 stocks at market price
	totalStocks = decimal.New(50, 0)
	sellOrdersFullyBought, sellOrderPartiallyBought, qtyPartialBought, qtyUnableToBuy, err := ob.ProcessMarketOrder(orderbook.Buy, totalStocks)
	// here sellOrdersFullyBought should be non nil
	if sellOrdersFullyBought == nil {
		logger.Panicf("sellOrdersFullyBought shouldnt be nil")
	}
	// and the sellOrdersFullyBought should have one sell order
	if len(sellOrdersFullyBought) != 1 {
		logger.Panicf("sellOrdersFullyBought should have one entry")
	}
	// that one entry should be the sell order from person "B" at price 99
	// because 99 is the best market price
	if sellOrdersFullyBought[0].ID() != "B-1" {
		logger.Panicf("the only entry in the sellOrdersFullyBought should be with id B-1 but its %v", sellOrdersFullyBought[0].ID())
	}
	if !sellOrdersFullyBought[0].Price().Equal(decimal.New(99, 0)) {
		logger.Panicf("the only entry in the sellOrdersFullyBought should be with price 99 but its %v", sellOrdersFullyBought[0].Price())
	}
	// the second variable sellOrderPartiallyBought should be nil because there was a sell order with qty 50
	// and our buy order of 50 matched, ie not resulted in partial buy
	if sellOrderPartiallyBought != nil {
		logger.Panicf("sellOrderPartiallyBought should be nil for buy from person X")
	}
	if !qtyPartialBought.Equal(decimal.Zero) {
		logger.Panicf("qtyBought should be 50 but its %v", qtyPartialBought)
	}

	// we are able to buy everything in single order
	if !qtyUnableToBuy.Equal(decimal.Zero) {
		logger.Panicf("qtyUnableToBuy should be 0 but its %v", qtyUnableToBuy)
	}
	logger.Println("success")
}
