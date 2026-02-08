package errors

type BaseError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *BaseError) Error() string {
	return e.Message
}

func BadRequest(msg string) *BaseError {
	return &BaseError{
		StatusCode: 400,
		Code:       "BAD_REQUEST",
		Message:    msg,
	}
}

func Unauthorized(msg string) *BaseError {
	return &BaseError{
		StatusCode: 401,
		Code:       "UNAUTHORIZED",
		Message:    msg,
	}
}

func NotFound(msg string) *BaseError {
	return &BaseError{
		StatusCode: 404,
		Code:       "NOT_FOUND",
		Message:    msg,
	}
}

func Internal(msg string) *BaseError {
	return &BaseError{
		StatusCode: 500,
		Code:       "INTERNAL_ERROR",
		Message:    msg,
	}
}
