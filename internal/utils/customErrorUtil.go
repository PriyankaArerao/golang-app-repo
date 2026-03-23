package utils

import (

	// "braces.dev/errtrace"

	"appointment-booking/internal/dto/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if errRes, ok := err.(response.ErrorResponse); ok {

		if err := c.JSON(errRes.StatusCode, errRes); err != nil {
			log.Error("Failed to send JSON response", "error", err)
		}
		return
	}

	if e := c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": err.Error(),
	}); e != nil {
		log.Error("Failed to send JSON response", "error", e)
	}
}
