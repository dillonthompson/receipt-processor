package api

import (
	"encoding/json"
	"net/http"

	"github.com/dillonthompson/receipt-processor/models"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func ReceiptHandler(c echo.Context) error {
	if c.Request().Method != "POST" {
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
	receipt := models.Receipt{}
	json.NewDecoder(c.Request().Body).Decode(&receipt)
	id, err := receiptProcessor(receipt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, models.ReceiptResponse{Id: id.String()})
}

func PointsHandler(c echo.Context) error {
	if c.Request().Method != "GET" {
		return c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
	response := models.PointsResponse{}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	response.Points = Receipts[id]
	return c.JSON(http.StatusOK, response)
}
