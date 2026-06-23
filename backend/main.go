package main

import (
	"h26_07_practice/router"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info("not loaded", "error", err)
	}
	e := echo.New()
	e.GET("/unread-management", router.GetUnreadManagement)
	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("fail!", "error", err)
	}
}
