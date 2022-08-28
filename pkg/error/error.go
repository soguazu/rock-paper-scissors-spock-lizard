package error

import "fmt"

type Error struct {
	status  string
	message string
	data    interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v\n", struct {
		status  string
		message string
		data    interface{}
	}{status: e.status, message: e.message, data: e.data})
}
