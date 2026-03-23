package main

import (
	"appointment-booking/cmd/api"
	"appointment-booking/config"
	"appointment-booking/internal/models"
	"appointment-booking/internal/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "appointment-booking/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Coach{},
		&models.User{},
		&models.Availability{},
		&models.Booking{},
	)

	e := echo.New()

	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler
	e.Validator = utils.NewValidator()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api.BindRoutes(e)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
