package config

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var NewFiberConfig = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		msg := ""
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			msg = e.Message
		}

		log.Println(code)
		log.Println(msg)
		// Send custom error page
		err = ctx.Status(code).JSON(fiber.Map{
			"ok":    false,
			"error": msg,
		})
		if err != nil {
			// In case the SendFile fails
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Internal server error",
			})
		}

		// Return from handler
		return nil
	},
}
