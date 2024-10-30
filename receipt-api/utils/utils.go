package utils

import (
	"fmt"
	"log"
	"math"
	"time"
)

// check if alphanumeric
func IsAlphanumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

// extract the day from a date string.
func GetDay(date string) (int, error) {
	// Parse the date string to ensure it follows the "YYYY-MM-DD" format
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, fmt.Errorf("invalid date format: %s, expected format is YYYY-MM-DD", date)
	}

	// Extract the day from the parsed date
	day := parsedDate.Day()
	return day, nil
}

func IsFloatMultiple(x, y, tolerance float64) bool {
	if y == 0 {
		return false // Division by zero is not allowed
	}

	result := x / y
	return math.Abs(result-math.Round(result)) < tolerance
}

// logAndReturnError logs an error message and returns it as a wrapped error.
func LogAndReturnError(message string, err error) error {
	log.Printf("%s: %v", message, err)        // Log the error
	return fmt.Errorf("%s: %w", message, err) // Wrap and return the error
}

func IsBetween2ToBefore4PM(timeStr string) (bool, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false, fmt.Errorf("invalid time format '%s', expected 24-hour format, but found: %w", timeStr, err)
	}
	return t.Hour() >= 14 && t.Hour() < 16, nil
}
