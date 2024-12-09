package orders

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type OrderHandler struct {
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) Sell(ctx *fiber.Ctx) error {
	var order Order
	err := ctx.BodyParser(&order)
	if err != nil {
		return err
	}
	return h.orderService.Sell(ctx.Context(), &order)
}

func (h *OrderHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	details, err := h.orderService.GetOrderDetails(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.JSON(details)
}

func (h *OrderHandler) GetAllOrders(ctx *fiber.Ctx) error {
	orders, err := h.orderService.GetOrders(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(orders)
}
