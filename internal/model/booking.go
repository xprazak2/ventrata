package model

type BookingStatus int

const (
	RESERVED = iota
	CONFIRMED
)

type Booking struct {
	Id             string
	Status         int
	Currency       string
	PricePerUnit   uint64
	ProductId      string
	AvailabilityId string
	Units          []TicketUnit
}

type TicketUnit struct {
	Id     string
	Ticket *string
}

func (bk *Booking) IsConfirmed() bool {
	return bk.Status == CONFIRMED
}

func (bk *Booking) TotalPrice() uint64 {
	return uint64(len(bk.Units) * int(bk.PricePerUnit))
}
