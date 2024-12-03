package orders

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Sell(ctx *fiber.Ctx) error {
	var order Order
	err := ctx.BodyParser(&order)
	if err != nil {
		return err
	}
	return nil
}
