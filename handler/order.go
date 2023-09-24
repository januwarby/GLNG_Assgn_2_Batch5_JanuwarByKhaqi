package handler

import (
	"h8-assignment-2/dto"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

// CreateOrder godoc
// @Tags orders
// @Description Create Order Data
// @ID create-new-order
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequest true "request body json"
// @Success 201 {object} dto.NewOrderResponse
// @Router /orders [post]
func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.CreateOrder(newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetOrders godoc
// @Tags orders
// @Description Get Order with Item Data
// @ID get-orders-with-items
// @Produce json
// @Success 200 {object} dto.GetOrdersResponse
// @Router /orders [get]
func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	response, err := oh.OrderService.GetOrders()

	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateOrder godoc
// @Tags orders
// @Description Update Order Data By Id
// @ID update-order
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequest true "request body json"
// @Param orderId path int true "order's id"
// @Success 200 {object} dto.NewOrderResponse
// @Router /orders/{orderId} [put]
func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.UpdateOrder(orderId, newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteOrder godoc
// @Tags orders
// @Description Delete Order Data By Id
// @ID delete-order
// @Param orderId path int true "order's id"
// @Success 200 {object} dto.NewOrderResponse
// @Router /orders/{orderId} [delete]
func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
    var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

    response, err := oh.OrderService.DeleteOrder(orderId)
    if err != nil {
        ctx.AbortWithStatusJSON(err.Status(), err)
        return
    }

    ctx.JSON(response.StatusCode, response)
}



