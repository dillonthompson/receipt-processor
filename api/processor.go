package api

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/dillonthompson/receipt-processor/models"
	"github.com/google/uuid"
)

type ReceiptsStore struct {
	Receipts map[uuid.UUID]int
	Mu       sync.Mutex
}

var Receipts = ReceiptsStore{Receipts: make(map[uuid.UUID]int)}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

func receiptProcessor(receipt models.Receipt) (uuid.UUID, error) {
	// create new id
	id := uuid.New()
	var err error

	// calculate points
	points := 0
	points += calculateAlphaNumericPoints(receipt.Retailer)
	points += calculateTotalPoints(receipt.Total)
	points += calculateItemPoints(receipt.Items)
	datePoints, err := calculateDatePoints(receipt.PurchaseDate)
	if err != nil {
		return uuid.UUID{}, err
	}
	timePoints, err := calculateTimePoints(receipt.PurchaseTime)
	if err != nil {
		return uuid.UUID{}, err
	}
	points += datePoints + timePoints

	// store id and points
	Receipts.Mu.Lock()
	defer Receipts.Mu.Unlock()
	Receipts.Receipts[id] = points
	return id, nil
}

func calculateAlphaNumericPoints(retailer string) int {
	return len(nonAlphanumericRegex.ReplaceAllString(retailer, ""))
}

func calculateTotalPoints(total float64) int {
	if total <= 0 {
		return 0
	}
	points := 0
	if total == math.Trunc(total) {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}
	return points
}

func calculateItemPoints(items []models.Item) int {
	if len(items) == 0 {
		return 0
	}
	points := 0
	points += (len(items) / 2) * 5
	for _, item := range items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}
	return points
}

func calculateDatePoints(purchaseDate string) (int, error) {
	points := 0
	t, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if t.Day()%2 != 0 {
		points += 6
	}
	return points, nil
}

func calculateTimePoints(purchaseTime string) (int, error) {
	points := 0
	t, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0, err
	}
	after, _ := time.Parse("15:04", "14:00")
	before, _ := time.Parse("15:04", "16:00")
	if t.After(after) && t.Before(before) {
		points += 10
	}
	return points, nil
}
