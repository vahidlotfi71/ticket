
type validatorMsgs struct {
	FieldName string
	Message   string
}

type Validator struct {
	fieldName string
	value     any
	file      *multipart.FileHeader
	skipFlag  bool
	ValMsgs   []validatorMsgs
	Error     error
}
