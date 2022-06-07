package errors

import "fmt"

type NotImplementedError struct{}

func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("[Error] not implemented")
}
