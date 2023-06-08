package api

import (
	"encoding/json"
	"net/http"

	"github.com/dillonthompson/receipt-processor/models"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func ReceiptHandler(c echo.Context) error {
	receipt := models.Receipt{}
	json.NewDecoder(c.Request().Body).Decode(&receipt)
	var id uuid.UUID
	var err error
	go func() {
		id, err = receiptProcessor(receipt)
	}()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, models.ReceiptResponse{Id: id.String()})
}

func PointsHandler(c echo.Context) error {
	response := models.PointsResponse{}
	id, err := uuid.Parse(c.Param("id"))
	var points int
	var ok bool
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	go func() {
		Receipts.Mu.Lock()
		defer Receipts.Mu.Unlock()
		points, ok = Receipts.Receipts[id]
	}()
	if !ok {
		return c.String(http.StatusNotFound, "Receipt not found")
	}
	response.Points = points
	return c.JSON(http.StatusOK, response)
}
