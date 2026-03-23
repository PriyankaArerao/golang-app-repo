package handlers

import (
	"net/http"
	"time"

	"appointment-booking/config"
	"appointment-booking/internal/dto/response"
	"appointment-booking/internal/models"

	"github.com/labstack/echo/v4"
)

// GetAvailableSlots
// @Summary Get available slots for a coach
// @Description Fetches all available 30-minute booking slots for a given coach on a specific date.
// @Tags Slots
// @Accept json
// @Produce json
// @Param coach_id query int true "Coach ID"
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {array} string "Array of available slot timestamps in RFC3339 format"
// @Failure 400 {object} response.ErrorResponse "Invalid request or date format"
// @Failure 404 {object} response.ErrorResponse "Availability not found for the specified coach"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /users/slots [get]
func GetAvailableSlots(c echo.Context) error {
	coachID := c.QueryParam("coach_id")
	dateStr := c.QueryParam("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid date")
	}

	day := date.Weekday().String()

	var availability models.Availability
	err = config.DB.Where("coach_id = ? AND day = ?", coachID, day).
		First(&availability).Error

	if err != nil {
		return response.ErrorResponse{
			ErrorCode:  "AVAILABILITY_NOT_FOUND",
			Message:    "Availability not found for the specified coach and date",
			StatusCode: http.StatusNotFound,
		}
	}

	start, _ := time.Parse("15:04", availability.StartTime)
	end, _ := time.Parse("15:04", availability.EndTime)

	var slots []string

	current := time.Date(date.Year(), date.Month(), date.Day(),
		start.Hour(), start.Minute(), 0, 0, time.UTC)

	endTime := time.Date(date.Year(), date.Month(), date.Day(),
		end.Hour(), end.Minute(), 0, 0, time.UTC)

	for current.Before(endTime) {
		slots = append(slots, current.Format(time.RFC3339))
		current = current.Add(30 * time.Minute)
	}

	// Remove booked slots
	var bookings []models.Booking
	config.DB.Where("coach_id = ? AND DATE(start_time) = ?", coachID, date).
		Find(&bookings)

	bookedMap := map[string]bool{}
	for _, b := range bookings {
		bookedMap[b.StartTime.Format(time.RFC3339)] = true
	}

	var available []string
	for _, s := range slots {
		if !bookedMap[s] {
			available = append(available, s)
		}
	}

	return c.JSON(http.StatusOK, available)
}
