package view

import "github.com/xprazak2/ventrata/internal/model"

type BaseBooking struct {
	Id             string `json:"id"`
	Status         int    `json:"status"`
	ProductId      string `json:"product_id"`
	AvailabilityId string `json:"availability_id"`
}

type Booking struct {
	BaseBooking
	Units []TicketUnit `json:"units"`
}

type PricedBooking struct {
	BaseBooking
	Units    []PricedTicketUnit `json:"units"`
	Price    uint64             `json:"price"`
	Currency string             `json:"currency"`
}

type TicketUnit struct {
	Id     string  `json:"id"`
	Ticket *string `json:"ticket"`
}

type PricedTicketUnit struct {
	TicketUnit
	Price    uint64 `json:"price"`
	Currency string `json:"currency"`
}

func pricedTicketUnitView(ticketUnit model.TicketUnit, price uint64, currency string) PricedTicketUnit {
	tu := TicketUnit{
		Id:     ticketUnit.Id,
		Ticket: ticketUnit.Ticket,
	}

	return PricedTicketUnit{
		TicketUnit: tu,
		Price:      price,
		Currency:   currency,
	}
}

func ticketUnitView(ticketUnit model.TicketUnit) TicketUnit {
	return TicketUnit{
		Id:     ticketUnit.Id,
		Ticket: ticketUnit.Ticket,
	}
}

func pricedTicketUnitsView(units []model.TicketUnit, price uint64, currency string) []PricedTicketUnit {
	res := make([]PricedTicketUnit, 0, len(units))
	for _, unit := range units {
		res = append(res, pricedTicketUnitView(unit, price, currency))
	}
	return res
}

func ticketUnitsView(units []model.TicketUnit) []TicketUnit {
	res := make([]TicketUnit, 0, len(units))
	for _, unit := range units {
		res = append(res, ticketUnitView(unit))
	}
	return res
}

func BookingView(booking model.Booking, isPriced bool) interface{} {
	baseBooking := BaseBooking{
		Id:             booking.Id,
		Status:         booking.Status,
		ProductId:      booking.ProductId,
		AvailabilityId: booking.AvailabilityId,
	}

	if isPriced {
		return PricedBooking{
			BaseBooking: baseBooking,
			Currency:    booking.Currency,
			Price:       booking.TotalPrice(),
			Units:       pricedTicketUnitsView(booking.Units, booking.PricePerUnit, booking.Currency),
		}
	}

	return Booking{
		BaseBooking: baseBooking,
		Units:       ticketUnitsView(booking.Units),
	}
}
