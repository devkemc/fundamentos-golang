package common

import (
	"github.com/gofiber/fiber/v2"
)

func NewError(statusCode int, message string) error {
	return fiber.NewError(statusCode, message)
}
