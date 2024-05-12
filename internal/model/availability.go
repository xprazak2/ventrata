package model

import "time"

type AvailabilityStatus int

const (
	AVAILABLE AvailabilityStatus = iota
	SOLD_OUT
)

type Availability struct {
	Id        string
	LocalDate time.Time
	Vacancies uint64
	Price     uint64
	Currency  string
}

func (avail *Availability) IsAvailable() bool {
	return avail.Vacancies > 0
}

func (avail *Availability) Status() AvailabilityStatus {
	if avail.IsAvailable() {
		return AVAILABLE
	}
	return SOLD_OUT
}
