package app

type Code string

type Error struct {
	Status  int    `json:"-"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details"`
}

func NewError(status int, code Code, message string, details ...any) *Error {
	err := &Error{
		Status:  status,
		Code:    code,
		Message: message,
	}

	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// Error 实现error接口
func (e *Error) Error() string {
	return e.Message
}

// GetStatus 实现huma的GetStatus()接口
func (e *Error) GetStatus() int {
	return e.Status
}
