package api

import (
	"appointment-booking/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func BindRoutes(e *echo.Echo) {
	e.POST("/coaches/availability", handlers.CreateAvailability)
	e.GET("/users/slots", handlers.GetAvailableSlots)
	e.POST("/users/bookings", handlers.CreateBooking)
	e.GET("/users/bookings", handlers.GetUserBookings)
}
