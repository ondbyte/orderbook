package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Handler struct {
	Service *Service
}

func HandlePlaceOrder(ctx *gin.Context) {
	form := new(OrderDetails)
	err := ctx.MustBindWith(form, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusTeapot, map[string]string{"error": err.Error()})
		return
	}

}
