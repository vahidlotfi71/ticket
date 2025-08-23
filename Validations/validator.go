// ticket/Validations/validator.go
package Validations

import (
	"fmt"

	"github.com/vahidlotfi71/ticket/Rules" // مسیر Rules خودتان
)

type Validator struct {
	fieldName string
	value     interface{}
	errors    []string
}

func NewValidator() *Validator {
	return &Validator{errors: []string{}}
}

func (v *Validator) Field(name string, value interface{}) *Validator {
	v.fieldName = name
	v.value = value
	return v
}

func (v *Validator) Required() *Validator {
	if Rules.IsEmpty(v.value) { // فرض: Rule ساده که خالی بودن رو چک می‌کنه
		v.errors = append(v.errors, fmt.Sprintf("%s الزامی است", v.fieldName))
	}
	return v
}

func (v *Validator) Email() *Validator {
	if !Rules.IsEmail(v.value) {
		v.errors = append(v.errors, fmt.Sprintf("%s فرمت ایمیل معتبر ندارد", v.fieldName))
	}
	return v
}

func (v *Validator) LengthBetween(min, max int) *Validator {
	if !Rules.LengthBetween(v.value, min, max) {
		v.errors = append(v.errors,
			fmt.Sprintf("%s باید بین %d و %d کاراکتر باشد", v.fieldName, min, max))
	}
	return v
}

func (v *Validator) Errors() []string {
	return v.errors
}

func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}
