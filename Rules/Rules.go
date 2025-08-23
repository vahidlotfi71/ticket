package Rules

import (
	"fmt"
	"regexp"
	"strings"
)

// بررسی خالی بودن فیلد
func Required(value string) string {
	if strings.TrimSpace(value) == "" {
		return "فیلد الزامی است"
	}
	return ""
}

// بررسی فرمت ایمیل
func Email(value string) string {
	if value == "" {
		return "ایمیل الزامی است"
	}
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(value) {
		return "ایمیل معتبر نیست"
	}
	return ""
}

// بررسی طول رشته بین min و max
func LengthBetween(value string, min int, max int) string {
	l := len(value)
	if l < min || l > max {
		return "طول رشته باید بین " + itoa(min) + " و " + itoa(max) + " باشد"
	}
	return ""
}

// بررسی شماره موبایل (مثال ایران: 09xxxxxxxxx)
func Phone(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "شماره موبایل الزامی است"
	}

	// الگوی شماره موبایل ایران: 09xxxxxxxxx
	re := regexp.MustCompile(`^09\d{9}$`)
	if !re.MatchString(value) {
		return "شماره موبایل معتبر نیست"
	}

	return ""
}

// تبدیل عدد به رشته (معادل strconv.Itoa)
func itoa(num int) string {
	return strings.Trim(strings.ReplaceAll(strings.TrimSpace(strings.Trim(fmt.Sprint(num), "[]")), " ", ""), "[]")
}
