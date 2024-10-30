package errors

import (
	"fmt"
	"log"
)

// logAndReturnError logs an error message and returns it as a wrapped error.
func LogAndReturnError(message string, err error) error {
	log.Printf("%s: %v", message, err)        // Log the error
	return fmt.Errorf("%s: %w", message, err) // Wrap and return the error
}
