package errors

import "strings"

type SystemError struct {
	BusinessError
	Reason error `json:"reason"`
}

func System(reason error, message, code string) SystemError {
	return SystemError{
		Reason:        reason,
		BusinessError: Business(message, code),
	}
}

func (e SystemError) ToJSON() []byte {
	return e.BusinessError.ToJSON()
}

func (e SystemError) WithErr(err error) SystemError {
	e.Reason = err
	return e
}

func (e SystemError) Error() string {
	var b strings.Builder

	b.WriteString(e.BusinessError.Error())
	b.WriteString(" (")
	b.WriteString(e.Reason.Error())
	b.WriteString(")")

	return b.String()
}
