# Go orderbook

Improved matching engine written in Go (Golang)

[![Go Report Card](https://goreportcard.com/badge/github.com/i25959341/orderbook)](https://goreportcard.com/report/github.com/i25959341/orderbook)
[![GoDoc](https://godoc.org/github.com/i25959341/orderbook?status.svg)](https://godoc.org/github.com/i25959341/orderbook)
[![gocover.run](https://gocover.run/github.com/i25959341/orderbook.svg?style=flat&tag=1.10)](https://gocover.run?tag=1.10&repo=github.com%2Fi25959341%2Forderbook)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)
[![Build Status](https://travis-ci.org/i25959341/orderbook.svg?branch=master)](https://travis-ci.org/i25959341/orderbook)

## Features
- Standard price-time priority
- Supports both market and limit orders
- Supports order cancelling
- High performance (above 300k trades per second)
- Optimal memory usage
- JSON Marshalling and Unmarsalling
- Calculating market price for definite quantity

## importing it in your project
```go
import (
	"github.com/ondbyte/orderbook"
)

func main(){
	ob := orderbook.NewOrderBook()
}
```
# EXAMPLE
## how to run the example
`cd <project root>`

`cd example`

run the program

`go run .`

program exits with `success` log

## Usage

example program where a order book begins with 
user A putting a limit sell order of 100 stocks at price 100
user B putting a limit sell order of 50 stocks at price 90

user X putting a market buy order of 50 stocks at market price

program verifies certain scenerios you should read the code to understand what is happening

```go
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


```

Following methods can be utilized:

```go

func (ob *OrderBook) ProcessLimitOrder(side Side, orderID string, quantity, price decimal.Decimal) (done []*Order, partial *Order, err error) { ... }

func (ob *OrderBook) ProcessMarketOrder(side Side, quantity decimal.Decimal) (done []*Order, partial *Order, quantityLeft decimal.Decimal, err error) { .. }

func (ob *OrderBook) CancelOrder(orderID string) *Order { ... }

```

## About primary functions

### ProcessLimitOrder

```go
// ProcessLimitOrder places new order to the OrderBook
// Arguments:
//      side     - what do you want to do (ob.Sell or ob.Buy)
//      orderID  - unique order ID in depth
//      quantity - how much quantity you want to sell or buy
//      price    - no more expensive (or cheaper) this price
//      * to create new decimal number you should use decimal.New() func
//        read more at https://github.com/shopspring/decimal
// Return:
//      error   - not nil if quantity (or price) is less or equal 0. Or if order with given ID is exists
//      done    - not nil if your order produces ends of anoter order, this order will add to
//                the "done" slice. If your order have done too, it will be places to this array too
//      partial - not nil if your order has done but top order is not fully done. Or if your order is
//                partial done and placed to the orderbook without full quantity - partial will contain
//                your order with quantity to left
//      partialQuantityProcessed - if partial order is not nil this result contains processed quatity from partial order
func (ob *OrderBook) ProcessLimitOrder(side Side, orderID string, quantity, price decimal.Decimal) (done []*Order, partial *Order, err error) { ... }
```

For example:
```
ProcessLimitOrder(ob.Sell, "uinqueID", decimal.New(55, 0), decimal.New(100, 0))

asks: 110 -> 5      110 -> 5
      100 -> 1      100 -> 56
--------------  ->  --------------
bids: 90  -> 5      90  -> 5
      80  -> 1      80  -> 1

done    - nil
partial - nil

```

```
ProcessLimitOrder(ob.Buy, "uinqueID", decimal.New(7, 0), decimal.New(120, 0))

asks: 110 -> 5
      100 -> 1
--------------  ->  --------------
bids: 90  -> 5      120 -> 1
      80  -> 1      90  -> 5
                    80  -> 1

done    - 2 (or more orders)
partial - uinqueID order

```

```
ProcessLimitOrder(ob.Buy, "uinqueID", decimal.New(3, 0), decimal.New(120, 0))

asks: 110 -> 5
      100 -> 1      110 -> 3
--------------  ->  --------------
bids: 90  -> 5      90  -> 5
      80  -> 1      90  -> 5

done    - 1 order with 100 price, (may be also few orders with 110 price) + uinqueID order
partial - 1 order with price 110

```

### ProcessMarketOrder

```go
// ProcessMarketOrder immediately gets definite quantity from the order book with market price
// Arguments:
//      side     - what do you want to do (ob.Sell or ob.Buy)
//      quantity - how much quantity you want to sell or buy
//      * to create new decimal number you should use decimal.New() func
//        read more at https://github.com/shopspring/decimal
// Return:
//      error        - not nil if price is less or equal 0
//      done         - not nil if your market order produces ends of anoter orders, this order will add to
//                     the "done" slice
//      partial      - not nil if your order has done but top order is not fully done
//      partialQuantityProcessed - if partial order is not nil this result contains processed quatity from partial order
//      quantityLeft - more than zero if it is not enought orders to process all quantity
func (ob *OrderBook) ProcessMarketOrder(side Side, quantity decimal.Decimal) (done []*Order, partial *Order, quantityLeft decimal.Decimal, err error) { .. }
```

For example:
```
ProcessMarketOrder(ob.Sell, decimal.New(6, 0))

asks: 110 -> 5      110 -> 5
      100 -> 1      100 -> 1
--------------  ->  --------------
bids: 90  -> 5      80 -> 1
      80  -> 2

done         - 2 (or more orders)
partial      - 1 order with price 80
quantityLeft - 0

```

```
ProcessMarketOrder(ob.Buy, decimal.New(10, 0))

asks: 110 -> 5
      100 -> 1
--------------  ->  --------------
bids: 90  -> 5      90  -> 5
      80  -> 1      80  -> 1
                    
done         - 2 (or more orders)
partial      - nil
quantityLeft - 4

```

### CancelOrder

```go
// CancelOrder removes order with given ID from the order book
func (ob *OrderBook) CancelOrder(orderID string) *Order { ... }
```

```
CancelOrder("myUinqueID-Sell-1-with-100")

asks: 110 -> 5
      100 -> 1      110 -> 5
--------------  ->  --------------
bids: 90  -> 5      90  -> 5
      80  -> 1      80  -> 1
                    
done         - 2 (or more orders)
partial      - nil
quantityLeft - 4

```

## License

The MIT License (MIT)

See LICENSE and AUTHORS files
