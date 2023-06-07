package api

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/dillonthompson/receipt-processor/models"
	"github.com/google/uuid"
)

var Receipts = map[uuid.UUID]int{}
var nonAlphanumericRegex = regexp.MustCompile(`[[:^alpha:]]`)

func receiptProcessor(receipt models.Receipt) (uuid.UUID, error) {
	// create new id
	id := uuid.New()
	// calculate points
	points := 0
	points += len(nonAlphanumericRegex.ReplaceAllString(receipt.Retailer, ""))
	fmt.Println(points)
	if receipt.Total == math.Trunc(receipt.Total) {
		points += 50
	}
	fmt.Println(points)
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}
	fmt.Println(points)
	points += (len(receipt.Items) / 2) * 5
	fmt.Println(points)
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}
	fmt.Println(points)
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}
	if date.Day()%2 != 0 {
		points += 6
	}
	fmt.Println(points)
	t, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}
	after, _ := time.Parse("15:04", "14:00")
	before, _ := time.Parse("15:04", "16:00")
	if t.After(after) && t.Before(before) {
		points += 10
	}
	fmt.Println(points)

	// store id and points
	Receipts[id] = points
	fmt.Println(Receipts)
	return id, nil
}
