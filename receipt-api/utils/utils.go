package utils

import (
	"math"
	"time"
)

// check if alphanumeric
func IsAlphanumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

// // extract the day from a date string.
// func GetDay(date string) (int, error) {
// 	// Parse the date string to ensure it follows the "YYYY-MM-DD" format
// 	parsedDate, err := time.Parse("2006-01-02", date)
// 	if err != nil {
// 		return 0, fmt.Errorf("invalid date format: %s, expected format is YYYY-MM-DD", date)
// 	}

// 	// Extract the day from the parsed date
// 	day := parsedDate.Day()
// 	return day, nil
// }

// extract the day from a date string.
func GetDay(date string) int {
	// date error was validated by the request validator
	parsedDate, _ := time.Parse("2006-01-02", date)

	// Extract the day from the parsed date
	return parsedDate.Day()

}

func IsFloatMultiple(x, y, tolerance float64) bool {
	if y == 0 {
		return false // Division by zero is not allowed
	}

	result := x / y
	return math.Abs(result-math.Round(result)) < tolerance
}

func IsBetween2ToBefore4PM(timeStr string) bool {
	// time has been validated before calling this function
	t, _ := time.Parse("15:04", timeStr)
	// if err != nil {
	// 	return false, fmt.Errorf("invalid time format '%s', expected 24-hour format, but found: %w", timeStr, err)
	// }
	return t.Hour() >= 14 && t.Hour() < 16
}
