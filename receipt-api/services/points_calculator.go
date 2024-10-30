package services

import (
	"math"
	"receipt-api/models"
	"receipt-api/utils"
	"strconv"
	"strings"
)

func CalculatePoints(receipt models.Receipt) (int, error) {
	points := 0

	//1 point for every alphanumeric character in the retailer name
	for _, ch := range receipt.Retailer {
		if utils.IsAlphanumeric(ch) {
			points++
		}
	}
	// the total is validated, no need to capture err here
	total_float_value, _ := strconv.ParseFloat(receipt.Total, 64)

	//50 points if the total is a round dollar amount with no cents
	if math.Trunc(total_float_value) == total_float_value {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if utils.IsFloatMultiple(total_float_value, 0.25, 0.00001) {
		points += 25
	}

	//5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	/**If the trimmed length of the item description is a multiple of 3,
		multiply the price by 0.2
		and round up to the nearest integer.
	The result is the number of points earned.*/

	for _, item := range receipt.Items {
		description_length := len(strings.TrimSpace(item.ShortDescription))

		if description_length%3 == 0 {
			// item price is validated, no need to capture parsing error
			// parse the price of the item
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	//6 points if the day in the purchase date is odd
	day := utils.GetDay(receipt.PurchaseDate)

	if day%2 != 0 {
		points += 6
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm
	isInTimeRange := utils.IsBetween2ToBefore4PM(receipt.PurchaseTime)

	if isInTimeRange {
		points += 10
	}

	return points, nil
}
