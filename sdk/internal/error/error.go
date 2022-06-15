package error

import (
	"fmt"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("API errors: %s", err.ErrorMessage)
}
