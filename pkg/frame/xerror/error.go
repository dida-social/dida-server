package xerror

type XError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewXError(code int, msg string) *XError {
	return &XError{
		Code: code,
		Msg:  msg,
	}
}

func (e *XError) Error() string {
	return e.Msg
}
