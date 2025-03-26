package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrosaraceno/errors/constants"
	"github.com/leandrosaraceno/errors/httperrors"
	"github.com/leandrosaraceno/errors/model"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		// fiber error
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(model.CustomError{
				StatusCode: fiberErr.Code,
				LogMessage: fiberErr.Message,
			})
		}

		// database error
		if dbErr := HandleDatabaseError(err); dbErr != nil {
			return c.Status(dbErr.StatusCode).JSON(dbErr)
		}

		if dbErr := HandleDatabaseError(err); dbErr != nil {
			return c.Status(dbErr.StatusCode).JSON(dbErr)
		}

		return httperrors.Handler(c, err)

		return c.Status(fiber.StatusInternalServerError).JSON(model.CustomError{
			StatusCode: fiber.StatusInternalServerError,
			LogMessage: err.Error(),
			Key:        constants.InternalServerError,
		})
	}
}
