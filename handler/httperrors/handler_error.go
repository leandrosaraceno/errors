package httperrors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/leandrosaraceno/errors/constants"
	"github.com/leandrosaraceno/errors/model"
	"github.com/leandrosaraceno/errors/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func HandlerHttpError(c *fiber.Ctx, err error) error {
	var gErr model.GenericError
	if errors.As(err, &gErr) {
		log.Error("HTTP Error<", strconv.Itoa(gErr.StatusCode), "> ", "Reason: ", err.Error())
		return c.Status(gErr.StatusCode).JSON(response.ReturnError(err.Error()))
	}

	log.Error("HTTP Error<", strconv.Itoa(http.StatusInternalServerError), "> ", "Reason: ", err.Error())
	return c.Status(http.StatusInternalServerError).JSON(response.ReturnError(constants.InternalServerError))
}
