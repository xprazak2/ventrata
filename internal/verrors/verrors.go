package verrors

import "fmt"

var (
	ErrInvalidDateRange     = fmt.Errorf("invalid date range")
	ErrInvalidRequestParams = fmt.Errorf("invalid request parameters")
	ErrNoBooking            = fmt.Errorf("at least one unit required for a booking to be created")
	ErrNoVacancies          = fmt.Errorf("not enough vacancies")
)
