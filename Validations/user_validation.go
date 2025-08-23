package Validations

import "github.com/vahidlotfi71/ticket/Models"

func UserValidation(user Models.User) *Validator {
	v := NewValidator()

	v.Field("name", user.Name).Required().LengthBetween(3, 50)
	v.Field("email", user.Email).Required().Email()
	v.Field("password", user.Password).Required().LengthBetween(6, 20)
	v.Field("phoneNumber", user.Phone).Required().Phone()

	return v
}
