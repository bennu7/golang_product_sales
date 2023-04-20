package responses

type CustomError struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	ErrorInfo  interface{} `json:"error_info,omitempty"`
}

func ErrResNotFound(errorInfo string) *CustomError {
	return &CustomError{
		StatusCode: 404,
		Message:    "Not Found",
		ErrorInfo:  errorInfo,
	}
}

func ErrResGeneral() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
}

func ErrBadrequest(message string, errorInfo interface{}) *CustomError {
	return &CustomError{
		StatusCode: 400,
		Message:    message,
		ErrorInfo:  errorInfo,
	}
}

func ErrCustom(statusCode int, message string, errorInfo interface{}) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
		ErrorInfo:  errorInfo,
	}
}
