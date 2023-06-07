package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dillonthompson/receipt-processor/models"
	"github.com/google/uuid"
)

func RequestHandler() {
	http.HandleFunc("/receipts/process", ReceiptHandler)
	http.HandleFunc("/receipts/{id}/points", PointsHandler)
}

func ReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	receipt := models.Receipt{}
	json.NewDecoder(r.Body).Decode(&receipt)
	id, err := receiptProcessor(receipt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ReceiptResponse{Id: id.String()})
}

func PointsHandler(w http.ResponseWriter, r *http.Request) {
	response := models.PointsResponse{}
	id, _ := uuid.Parse(strings.Split(r.URL.Path, "/")[1])
	response.Points = Receipts[id]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
