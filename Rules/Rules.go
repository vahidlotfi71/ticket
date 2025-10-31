package rules

import (
	"fmt"
	"strings"
)

func Required(value string) string {
	if strings.TrimSpace(value) == "" {
		return "The filed is Required"
	}
	return ""
}

func LengThBetween(value string, min, max int) string {
	l := len(value)
	if l < min || l > max {
		return fmt.Sprintf("length must be between %d and %d itoa(max)", min, max)
	}
	return ""
}
