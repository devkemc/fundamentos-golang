package orders

import "github.com/gofiber/fiber/v2"

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
	return nil
}

func (h *OrderHandler) GetAllOrders(ctx *fiber.Ctx) error {
	return nil
}
