package handlers

import (
	"net/http"

	"appointment-booking/config"
	"appointment-booking/internal/dto/request"
	"appointment-booking/internal/dto/response"
	"appointment-booking/internal/models"

	"github.com/labstack/echo/v4"
)

// CreateAvailability godoc
// @Summary Create coach availability
// @Description Allows a coach to set their availability for a day of the week
// @Tags Coaches
// @Accept json
// @Produce json
// @Param data body request.CreateAvailabilityRequest true "Availability request"
// @Success 200 {object} request.CreateAvailabilityRequest
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /coaches/availability [post]
func CreateAvailability(c echo.Context) error {
	var req request.CreateAvailabilityRequest

	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse{
			ErrorCode:  "INVALID_REQUEST",
			Message:    "Invalid request payload",
			StatusCode: http.StatusBadRequest,
		}
	}

	if err := c.Validate(&req); err != nil {
		return response.ErrorResponse{
			ErrorCode:  "VALIDATION_ERROR",
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
	}

	availability := models.Availability{
		CoachID:   req.CoachID,
		Day:       req.Day,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	if err := config.DB.Create(&availability).Error; err != nil {
		return response.ErrorResponse{
			ErrorCode:  "INTERNAL_SERVER_ERROR",
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return c.JSON(http.StatusOK, req)
}
