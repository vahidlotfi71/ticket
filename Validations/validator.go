package Validations

import "github.com/vahidlotfi71/ticket/Rules"

type Validator struct {
	fieldName string
	value     string
	errors    map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		errors: make(map[string]string),
	}
}

// تعیین فیلد و مقدار آن
func (v *Validator) Field(name string, value string) *Validator {
	v.fieldName = name
	v.value = value
	return v
}

// بررسی الزامی بودن فیلد
func (v *Validator) Required() *Validator {
	if err := Rules.Required(v.value); err != "" {
		v.errors[v.fieldName] = err
	}
	return v
}

// بررسی طول رشته
func (v *Validator) LengthBetween(min, max int) *Validator {
	if err := Rules.LengthBetween(v.value, min, max); err != "" {
		v.errors[v.fieldName] = err
	}
	return v
}

// بررسی ایمیل
func (v *Validator) Email() *Validator {
	if err := Rules.Email(v.value); err != "" {
		v.errors[v.fieldName] = err
	}
	return v
}
func (v *Validator) Phone() *Validator {
	if err := Rules.Phone(v.value); err != "" {
		v.errors[v.fieldName] = err
	}
	return v
}

// دریافت همه خطاها
func (v *Validator) Errors() map[string]string {
	return v.errors
}

// آیا داده‌ها معتبر هستند؟
func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}
