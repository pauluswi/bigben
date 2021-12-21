package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pauluswi/bigben/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
