package validation

import "vahid/Models"

func UserValidation(user Models.Users) *Validator {
	v := NewValidator()

	v.Field("name", user.Name).Required().LengthBetween(10, 20)
	v.Field("email", user.Email).Required().LengthBetween(2, 5)
	v.Field("password", user.Password).Required().LengthBetween(8, 11)
	v.Field("phone", user.Phone).Required().LengthBetween(9, 9)

	return v
}
