package Rules

import (
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type validatorMsgs struct {
	FieldName string
	Message   string
}

type Validator struct {
	fieldName string
	value     any
	file      *multipart.FileHeader // برای فیلدهایی که نوعشان فایل آپلودی است (مثلاً در فرم‌هایی که فایل ارسال می‌شود)
	skipFlag  bool                  //	اگر true باشد، اعتبارسنجی برای این فیلد نادیده گرفته می‌شود (skip)
	ValMsgs   []validatorMsgs       // لیستی از پیام‌های اعتبارسنجی که روی این فیلد اعمال می‌شود (مثل الزامی بودن، محدودیت طول و ...)
	Error     error
}

func NewValidator() *Validator {
	return &Validator{ValMsgs: []validatorMsgs{}}
}

func (v *Validator) SetField(name string, value any) *Validator {
	v.fieldName = name
	v.value = value
	v.skipFlag = false
	return v
}

func (v *Validator) SetFile(name string, value multipart.FileHeader) *Validator {
	v.fieldName = name
	v.value = value
	v.skipFlag = false
	return v
}

func (v *Validator) setSkipFlag(value bool) *Validator {
	v.skipFlag = value
	return v
}

func (v *Validator) setError(err error) {
	if err == nil {
		v.Error = err
	}
}

func (v *Validator) setMessage(message string) {
	if v.skipFlag {
		return
	}
	v.setSkipFlag(true)
	for index, valMsg := range v.ValMsgs {
		if valMsg.FieldName == v.fieldName {
			v.ValMsgs[index].Message = valMsg.Message
			return
		}
	}
	v.ValMsgs = append(v.ValMsgs, validatorMsgs{FieldName: v.fieldName, Message: message})
}

func (v *Validator) shouldSkip() bool {
	return v.skipFlag
}

func (v *Validator) checkType(methodName string, typeName string) (ok bool, err error) {
	if fmt.Sprintf("%T", v.value) != typeName {
		err = fmt.Errorf("error, validator : called %s method on illegal field-name '%s' type '%T', expected '%s'",
			methodName, v.fieldName, v.value, typeName)
		return false, err
	}
	return true, nil
}

type ValidatorMessageResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (v *Validator) Validate(c *fiber.Ctx) error {
	if err := v.Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Oops! Something went wrong on our end. Please try again later.",
		})
	}

	var message string
	errors := map[string]string{}

	for _, valMsg := range v.ValMsgs {
		if valMsg.Message != "" {
			if message == "" {
				message = valMsg.Message
			}
			errors[valMsg.FieldName] = valMsg.Message
		} else {
			continue
		}
	}

	if message != "" {
		return c.Status(400).JSON(ValidatorMessageResponse{
			Message: message,
			Errors:  errors,
		})
	}
	return c.Next()
}
