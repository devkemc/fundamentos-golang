package orders

import "github.com/gofiber/fiber/v2"

func SetupRoutes(group fiber.Router, handler *OrderHandler) {
	orders := group.Group("/orders")
	{
		orders.Post("", handler.Sell)
		orders.Get("/:id", handler.GetOrderDetails)
		orders.Get("", handler.GetAllOrders)
	}
}
