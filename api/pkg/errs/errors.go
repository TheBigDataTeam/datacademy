package errs

type InternalError struct {
	code uint64
	msg  string
	desc string
	data interface{}
}

func (ie InternalError) Error() string {
	return ie.msg + ": " + ie.desc
}

func (ie InternalError) Code() uint64 {
	return ie.code
}

func (ie InternalError) Message() string {
	return ie.msg
}

func (ie InternalError) Description() string {
	return ie.desc
}

func OperationFailed(code uint64, msg, desc string, data interface{}) InternalError {
	return InternalError{
		code: code,
		msg:  msg,
		desc: desc,
		data: data,
	}
}
