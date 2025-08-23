// ticket/Rules/rules.go
package Rules

import (
	"regexp"
	"strings"
)

func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v) == ""
	}
	return false
}

func IsEmail(value interface{}) bool {
	if str, ok := value.(string); ok {
		re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
		return re.MatchString(str)
	}
	return false
}

func LengthBetween(value interface{}, min, max int) bool {
	if str, ok := value.(string); ok {
		l := len(str)
		return l >= min && l <= max
	}
	return false
}
