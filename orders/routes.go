package orders

import "github.com/gofiber/fiber/v2"

func SetupRoutes(group fiber.Router) {
	orderHandler := NewHandler()
	orders := group.Group("/orders")
	{
		orders.Post("", orderHandler.Sell)
	}
}
