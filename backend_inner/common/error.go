package common

type Errors struct {
	statusCode int
	msg        string
}

func (errors *Errors) NewError(code int, errMsg string) Errors {
	return Errors{
		statusCode: code,
		msg:        errMsg,
	}
}

func (errors *Errors) GetError() (int, string) {
	return errors.statusCode, errors.msg
}
