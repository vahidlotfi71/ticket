package rules

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// قانون الزامی: بررسی می‌کند مقدار رشته خالی نباشد
func (v *Validator) Required(message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || strings.TrimSpace(str) == "" {
		v.setMessage(message)
	}
	return v
}

// قانون حداقل طول: بررسی می‌کند رشته حداقل طول مشخصی داشته باشد
func (v *Validator) MinLength(length int, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || utf8.RuneCountInString(str) < length {
		v.setMessage(message)
	}
	return v
}

// قانون حداکثر طول: بررسی می‌کند رشته حداکثر طول مشخصی داشته باشد
func (v *Validator) MaxLength(length int, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || utf8.RuneCountInString(str) > length {
		v.setMessage(message)
	}
	return v
}

// قانون طول دقیق: بررسی می‌کند رشته دقیقاً طول مشخصی داشته باشد
func (v *Validator) ExactLength(length int, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || utf8.RuneCountInString(str) != length {
		v.setMessage(message)
	}
	return v
}

// قانون فرمت ایمیل: بررسی می‌کند رشته فرمت ایمیل معتبر داشته باشد
func (v *Validator) IsEmail(message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || !isValidEmail(str) {
		v.setMessage(message)
	}
	return v
}

// قانون تطابق با الگو: بررسی می‌کند رشته با الگوی regex تطابق داشته باشد
func (v *Validator) Regex(pattern string, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.setMessage(message)
		return v
	}

	matched, err := regexp.MatchString(pattern, str)
	if err != nil || !matched {
		v.setMessage(message)
	}
	return v
}

// قانون شامل: بررسی می‌کند رشته حاوی زیررشته مشخص باشد
func (v *Validator) Contains(substr string, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || !strings.Contains(str, substr) {
		v.setMessage(message)
	}
	return v
}

// قانون شروع با: بررسی می‌کند رشته با پیشوند مشخص شروع شود
func (v *Validator) StartsWith(prefix string, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || !strings.HasPrefix(str, prefix) {
		v.setMessage(message)
	}
	return v
}

// قانون پایان با: بررسی می‌کند رشته با پسوند مشخص پایان یابد
func (v *Validator) EndsWith(suffix string, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok || !strings.HasSuffix(str, suffix) {
		v.setMessage(message)
	}
	return v
}

// قانون وجود در لیست: بررسی می‌کند رشته در لیست مقادیر مجاز وجود داشته باشد
func (v *Validator) In(allowed []string, message string) *Validator {
	if v.shouldSkip() {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.setMessage(message)
		return v
	}

	found := false
	for _, val := range allowed {
		if val == str {
			found = true
			break
		}
	}

	if !found {
		v.setMessage(message)
	}
	return v
}

// تابع کمکی برای اعتبارسنجی ایمیل
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
