package static

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewOrderPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "new order",
			`
			<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>New Order Form</title>
    <!-- Bootstrap CSS -->
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
</head>

<body>

    <div class="container">
        <form class="form-horizontal" action="https://example.com/submit" method="post">
            <fieldset>

                <!-- Form Name -->
                <legend>New Order</legend>

                <!-- Select Basic -->
                <div class="form-group">
                    <label class="col-md-4 control-label" for="selectbasic">Order Type</label>
                    <div class="col-md-4">
                        <select id="selectbasic" name="selectbasic" class="form-control">
                            <option value="buy">Buy</option>
                            <option value="sell">Sell</option>
                        </select>
                    </div>
                </div>

                <!-- Text input-->
                <div class="form-group">
                    <label class="col-md-4 control-label" for="quantity">Quantity</label>
                    <div class="col-md-4">
                        <input id="quantity" name="quantity" type="text" placeholder="Quantity"
                            class="form-control input-md" required="">

                    </div>
                </div>

                <!-- Text input-->
                <div class="form-group">
                    <label class="col-md-4 control-label" for="price">Price</label>
                    <div class="col-md-4">
                        <input id="price" name="price" type="text" placeholder="Price" class="form-control input-md">
                        <span class="help-block">If you enter the price, the order will be a limit order. Otherwise,
                            it will be a market order.</span>
                    </div>
                </div>

                <!-- Button -->
                <div class="form-group">
                    <label class="col-md-4 control-label" for="place_order"></label>
                    <div class="col-md-4">
                        <button id="place_order" name="place_order" class="btn btn-success">Place Order</button>
                    </div>
                </div>

            </fieldset>
        </form>
    </div>

    <!-- Bootstrap JS and jQuery -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>

</body>

</html>

			`,
		)
	}
}
