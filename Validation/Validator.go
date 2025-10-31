package validation

import (
	rules "vahid/Rules"
)

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

func (v *Validator) Field(name string, value string) *Validator {
	v.fieldName = name
	v.value = value
	return v
}

func (v *Validator) Required() *Validator {
	if err := rules.Required(v.value); err != "" {
		v.errors[v.fieldName] = err
	}
	return v
}

func (v *Validator) LengthBetween(min, max int) *Validator {
	if err := rules.LengThBetween(v.value, min, max); err != "" {
		v.errors[v.fieldName] = err
	}
	return v
}
