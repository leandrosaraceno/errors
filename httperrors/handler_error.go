package httperrors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrosaraceno/errors/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Handler(c *fiber.Ctx, err error) error {
	var status int
	switch err.(type) {
	case BadRequestError:
		status = http.StatusBadRequest
	case UnauthorizedError:
		status = http.StatusUnauthorized
	case NotFoundError:
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}

	log.Error("HTTP Error<", strconv.Itoa(status), "> ", "Reason: ", err.Error())
	return c.Status(status).JSON(response.ReturnError(err.Error()))
}
