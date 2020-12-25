package models

type Result struct {
	Code	int
	Message	string
	Data	interface{}
}

const (
	SUCCESS = iota
	FAIL
)

func ResponseResult(code int, message string, data interface{}) (result Result) {
	result.Code = code
	result.Message = message
	result.Data = data
	return
}
