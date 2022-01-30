package errs

type InternalError struct {
	code int
	msg  string
	desc string
	data interface{}
}

func (ie InternalError) Error() string {
	return ie.msg + ": " + ie.desc
}

func (ie InternalError) Code() int {
	return ie.code
}

func (ie InternalError) Message() string {
	return ie.msg
}

func (ie InternalError) Description() string {
	return ie.desc
}

func OperationFailed(code int, msg, desc string, data interface{}) *InternalError {
	return &InternalError{
		code: code,
		msg:  msg,
		desc: desc,
		data: data,
	}
}
