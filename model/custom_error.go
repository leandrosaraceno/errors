package model

type CustomError struct {
	Key        string `json:"error_key"`
	LogMessage string
	StatusCode int `json:"status_code"`
}

func (e *CustomError) Error() string {
	return e.Key
}

func NewCustomError(key string, statusCode int) *CustomError {
	return &CustomError{
		Key:        key,
		StatusCode: statusCode,
	}
}
