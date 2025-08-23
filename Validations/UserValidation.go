// ticket/Validations/user_validation.go
package Validations

import "github.com/vahidlotfi71/ticket/Models"

func UserValidation(user Models.User) *Validator {
	v := NewValidator()

	v.Field("Username", user.Username).Required().LengthBetween(3, 20)
	v.Field("Email", user.Email).Required().Email()
	v.Field("Password", user.Password).Required().LengthBetween(6, 32)

	return v
}
