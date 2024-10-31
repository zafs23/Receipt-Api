package validators

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
)

// validate receipt fields
func ValidateReceipt(receipt models.Receipt) error {

	// check if the fields are not empty
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		return errors.New("missing required fields in the request")
	}

	// check if the receipt fields have correct format

	// Date format validation (YYYY-MM-DD)
	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return fmt.Errorf("invalid date format for PurchaseDate: %v, expected format is YYYY-MM-DD", err)
	}

	// Time format validation (HH:MM)
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return fmt.Errorf("invalid time format for PurchaseTime: %v, expected 24-hour format HH:MM", err)
	}

	// Total validation: must be a non-negative decimal number with up to two decimal places
	if err := validateDecimal(receipt.Total); err != nil {
		return fmt.Errorf("invalid format for Total: %v", err)
	}

	// Validate each item in the Items array
	for _, item := range receipt.Items {
		// Ensure ShortDescription has a reasonable length
		if len(strings.TrimSpace(item.ShortDescription)) == 0 || len(item.ShortDescription) > 50 {
			return errors.New("item ShortDescription must be between 1 and 50 characters")
		}

		// Validate Price format and non-negative value
		if err := validateDecimal(item.Price); err != nil {
			return fmt.Errorf("invalid format for item Price: %v", err)
		}
	}

	return nil

}

// validateDecimal checks if a string represents a non-negative decimal number with up to two decimal places.
func validateDecimal(value string) error {
	// Check if the value is a valid number
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil || parsedValue < 0 {
		return errors.New("value must be a non-negative decimal number")
	}

	// Ensure it has at most two decimal places
	matched, _ := regexp.MatchString(`^\d+(\.\d{1,2})?$`, value)
	if !matched {
		return errors.New("value must have at most two decimal places")
	}

	return nil
}
