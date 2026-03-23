package handlers

import (
	"net/http"
	"time"

	"appointment-booking/config"
	"appointment-booking/internal/dto/request"
	"appointment-booking/internal/dto/response"
	"appointment-booking/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

// CreateBooking
// @Summary Create a booking
// @Description Allows a user to book a 30-minute slot with a coach. Prevents double booking.
// @Tags Bookings
// @Accept json
// @Produce json
// @Param data body request.CreateBookingRequest true "Booking request"
// @Success 200 {object} models.Booking "Booking created successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid request or validation error"
// @Failure 409 {object} response.ErrorResponse "Time slot already booked"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /users/bookings [post]
func CreateBooking(c echo.Context) error {

	var req request.CreateBookingRequest

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

	t, err := time.Parse(time.RFC3339, req.Time)
	if err != nil {
		return response.ErrorResponse{
			ErrorCode:  "INVALID_REQUEST",
			Message:    "Invalid datetime format",
			StatusCode: http.StatusBadRequest,
		}
	}

	booking := models.Booking{
		UserID:    req.UserID,
		CoachID:   req.CoachID,
		StartTime: t,
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		var existing models.Booking

		err := tx.Where("coach_id = ? AND start_time = ?", req.CoachID, t).
			First(&existing).Error

		if err == nil {
			return response.ErrorResponse{
				ErrorCode:  "SLOT_ALREADY_BOOKED",
				Message:    "This time slot is already booked",
				StatusCode: http.StatusConflict,
			}
		}

		if err != gorm.ErrRecordNotFound {
			log.Error("Database error while checking existing booking", "error", err)
			return response.ErrorResponse{

				ErrorCode:  "INTERNAL_SERVER_ERROR",
				Message:    "Internal server error",
				StatusCode: http.StatusInternalServerError,
			}
		}

		return tx.Create(&booking).Error
	})

	if err != nil {
		log.Error("Database error", "error", err)
		return response.ErrorResponse{

			ErrorCode:  "INTERNAL_SERVER_ERROR",
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return c.JSON(http.StatusOK, booking)
}

// GetUserBookings
// @Summary Get user bookings
// @Description Fetch all upcoming bookings for a specific user
// @Tags Bookings
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {array} response.BookingResponse "List of bookings"
// @Failure 400 {object} response.ErrorResponse "Invalid request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /users/bookings [get]
func GetUserBookings(c echo.Context) error {
	userID := c.QueryParam("user_id")

	var bookings []response.BookingResponse
	config.DB.Where("user_id = ?", userID).Find(&bookings)

	return c.JSON(http.StatusOK, bookings)
}
