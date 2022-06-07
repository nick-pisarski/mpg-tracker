package errors

import "fmt"

type NotFoundError struct {
	Id     int
	Action string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("[ERROR] %s failed, item: %d, not found", e.Action, e.Id)
}
