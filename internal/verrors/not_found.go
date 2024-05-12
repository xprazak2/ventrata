package verrors

import "fmt"

type NotFoundError struct {
	resource string
	id       string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with id '%s' not found", e.resource, e.id)
}

func NewNotFoundError(resource string, id string) NotFoundError {
	return NotFoundError{resource, id}
}
