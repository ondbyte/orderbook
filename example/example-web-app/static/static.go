package static

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewOrderPage(action string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		html := fmt.Sprintf(`
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
    <form class="form-horizontal" action="%v" method="post">
        <fieldset>
    
            <!-- Form Name -->
            <legend>New Order</legend>
    
            <!-- Select Basic -->
            <div class="form-group">
                <label class="col-md-4 control-label" for="side">Order Type</label>
                <div class="col-md-4">
                    <select id="side" name="side" class="form-control">
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
    
        `, action)
		ctx.Header("Content-Type", "text/html")
		ctx.Data(http.StatusOK, "charset=utf-8", []byte(html))
	}
}

func NewOrderUpdatesPage(action string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		html := fmt.Sprintf(`<!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>WebSocket JSON Data</title>
        </head>
        <body>
            <div id="data-container"></div>
        
            <script>
            // Get the current URL
                var currentURL = window.location.href;
                var parser = document.createElement('a');
                parser.href = currentURL;

                var domainWithPort = parser.hostname + (parser.port ? ':' + parser.port : '');
                const socketAddress = "ws://"+domainWithPort+"%v";
                // Establishing a WebSocket connection
                const socket = new WebSocket(socketAddress);
        
                // Handling WebSocket connection opened event
                socket.addEventListener("open", (event) => {
                    console.log("WebSocket Connection Established!");
                });
        
                // Handling WebSocket messages received event
                socket.addEventListener("message", (event) => {
                    // Parsing JSON data received from the server
                    const jsonData = JSON.parse(event.data);
        
                    // Displaying JSON data in the data-container div
                    const dataContainer = document.getElementById("data-container");
                    dataContainer.innerHTML = JSON.stringify(jsonData, null, 2);
                });
        
                // Handling WebSocket connection closed event
                socket.addEventListener("close", (event) => {
                    console.log("WebSocket Connection Closed!");
                });
        
                // Handling WebSocket errors
                socket.addEventListener("error", (event) => {
                    console.error("WebSocket Error: ", event);
                });
            </script>
        </body>
        </html>
        `, action)
		ctx.Header("Content-Type", "text/html")
		ctx.Data(http.StatusOK, "charset=utf-8", []byte(html))
	}
}
