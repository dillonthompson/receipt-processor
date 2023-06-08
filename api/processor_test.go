package api

import (
	"testing"

	"github.com/dillonthompson/receipt-processor/models"
)

func TestCalculateAlphaNumericPoints(t *testing.T) {
	tt := []struct {
		name     string
		retailer string
		want     int
	}{
		{"should get 7 with normal name", "Walmart", 7},
		{"should get 7 with nonalphanumeric", "Walmart!", 7},
		{"should get 10 with numbers", "Walmart123", 10},
		{"should get 10 with nonalphanumeric", "Walmart!123", 10},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateAlphaNumericPoints(tc.retailer)
			if got != tc.want {
				t.Errorf("calculateAlphaNumericPoints(%s) = %d; want %d", tc.retailer, got, tc.want)
			}
		})
	}
}

func TestCalculateTotalPoints(t *testing.T) {
	tt := []struct {
		name  string
		total float64
		want  int
	}{
		{"should get 75 with whole number", 100.00, 75},
		{"should get 25 with quarter", 100.25, 25},
		{"should get 0 with non quarter", 100.26, 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateTotalPoints(tc.total)
			if got != tc.want {
				t.Errorf("calculateTotalPoints(%f) = %d; want %d", tc.total, got, tc.want)
			}
		})
	}
}

func TestCalculateItemPoints(t *testing.T) {
	tt := []struct {
		name  string
		items []models.Item
		want  int
	}{
		{
			"should get 5 with 2 items and no description divisible by 3",
			[]models.Item{
				{
					ShortDescription: "item1",
					Price:            6.00,
				},
				{
					ShortDescription: "item2",
					Price:            6.00,
				},
			},
			5,
		},
		{
			"should get 2 with 1 item that has description divisible by 3 and a price of 6.00",
			[]models.Item{
				{
					ShortDescription: "item11",
					Price:            6.00,
				},
			},
			2,
		},
		{
			"should get 9 with 3 items and only one with description divisible by 3 and a price of 20.00",
			[]models.Item{
				{
					ShortDescription: "item1",
					Price:            6.00,
				},
				{
					ShortDescription: "item2",
					Price:            6.00,
				},
				{
					ShortDescription: "item33",
					Price:            20.00,
				},
			},
			9,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateItemPoints(tc.items)
			if got != tc.want {
				t.Errorf("calculateItemPoints(%v) = %d; want %d", tc.items, got, tc.want)
			}
		})
	}
}

func TestCalculateDatePoints(t *testing.T) {
	tt := []struct {
		name         string
		purchaseDate string
		want         int
	}{
		{"should get 6 with odd number day", "2020-05-05", 6},
		{"should get 0 with even number day", "2020-05-06", 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := calculateDatePoints(tc.purchaseDate)
			if got != tc.want {
				t.Errorf("calculateDatePoints(%s) = %d; want %d", tc.purchaseDate, got, tc.want)
			}
		})
	}
}

func TestCalulateTimePoints(t *testing.T) {
	tt := []struct {
		name         string
		purchaseTime string
		want         int
	}{
		{"should get 10 with time between 14:00 and 16:00", "15:00", 10},
		{"should get 0 with time not between 14:00 and 16:00", "16:01", 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := calculateTimePoints(tc.purchaseTime)
			if got != tc.want {
				t.Errorf("calculateTimePoints(%s) = %d; want %d", tc.purchaseTime, got, tc.want)
			}
		})
	}
}
