package main

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func countAlphanumeric(data receipt) int {
	count := 0

	for _, char := range data.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

func isRoundTotal(data receipt) int {
	total, err := strconv.ParseFloat(data.Total, 64)
	if err != nil {
		return 0
	}
	if total == math.Trunc(total) {
		return 50
	}
	return 0
}

func isMultipleOf025(data receipt) int {
	total, err := strconv.ParseFloat(data.Total, 64)
	if err != nil {
		return 0
	}
	if math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

func countItems(data receipt) int {
	return (len(data.Items) / 2) * 5
}

func pointsPerItem(data receipt) int {
	points := 0
	for _, item := range data.Items {
		if math.Mod(float64(len(strings.TrimSpace(item.Description))), 3.0) == 0.0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				additional_points := int(math.Ceil(price * 0.2))
				points += additional_points
			}
		}
	}
	return points
}

func isOddDay(data receipt) int {
	date, err := time.Parse("2006-01-02", data.PurchaseDate)
	if err != nil {
		return 0
	}
	if math.Mod(float64(date.Day()), 2.0) == 1.0 {
		return 6
	}
	return 0
}

func isBetween2And4PM(data receipt) int {
	date, err := time.Parse("15:04", data.PurchaseTime)
	if err != nil {
		return 0
	}
	hour := date.Hour()
	if hour >= 14 && hour < 16 {
		return 10
	}
	return 0
}

func calculateAllPoints(r receipt) int {
	return countAlphanumeric(r) + isRoundTotal(r) + isMultipleOf025(r) + countItems(r) + pointsPerItem(r) + isOddDay(r) + isBetween2And4PM(r)
}
