package api

import (
	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	e := echo.New()
	e.POST("/receipts/process", ReceiptHandler)
	e.GET("/receipts/:id/points", PointsHandler)
	return e
}
