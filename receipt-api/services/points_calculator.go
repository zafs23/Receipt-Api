package services

import (
	"fmt"
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
	// convert price total to float
	total_float_value, float_err := strconv.ParseFloat(receipt.Total, 64)
	if float_err != nil {
		return 0, utils.LogAndReturnError(fmt.Sprintf("error parsing total '%s' to float64", receipt.Total), float_err)
	}

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
			// parse the price of the item
			price, price_err := strconv.ParseFloat(item.Price, 64)
			if price_err != nil {
				//return 0, fmt.Errorf("error parsing price '%s' for item '%s': %w", item.Price, item.ShortDescription, err)
				return 0, utils.LogAndReturnError(fmt.Sprintf("error extracting price from purchase item '%s'", item.ShortDescription), price_err)
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	//6 points if the day in the purchase date is odd
	day, date_err := utils.GetDay(receipt.PurchaseDate)
	if date_err != nil {
		//return 0, fmt.Errorf("error extracting day from purchase date '%s': %w", receipt.PurchaseDate, date_err)
		return 0, utils.LogAndReturnError(fmt.Sprintf("error extracting day from purchase date '%s'", receipt.PurchaseDate), date_err)
	}

	if day%2 != 0 {
		points += 6
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm
	isInTimeRange, time_err := utils.IsBetween2ToBefore4PM(receipt.PurchaseTime)
	if time_err != nil {
		//return 0, fmt.Errorf("error extracting time from purchase time '%s': %w", receipt.PurchaseDate, time_err)
		return 0, utils.LogAndReturnError(fmt.Sprintf("error extracting time from purchase time '%s'", receipt.PurchaseTime), time_err)
	}

	if isInTimeRange {
		points += 10
	}

	return points, nil
}
