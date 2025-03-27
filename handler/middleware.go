package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrosaraceno/errors/handler/httperrors"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		// database error
		err = HandleDatabaseError(err)

		return httperrors.HandlerHttpError(c, err)

	}
}
