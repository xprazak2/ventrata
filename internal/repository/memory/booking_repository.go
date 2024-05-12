package memory

import (
	"github.com/google/uuid"
	"github.com/xprazak2/ventrata/internal/model"
)

func (repo *Memory) CreateBooking(
	productId string, availId string, units uint64, price uint64, currency string,
) (*model.Booking, error) {
	id := uuid.NewString()

	booking := model.Booking{
		Id:             id,
		Status:         model.RESERVED,
		ProductId:      productId,
		AvailabilityId: availId,
		Units:          createUnits(units),
		Currency:       currency,
		PricePerUnit:   price,
	}

	repo.bookings[id] = booking

	return &booking, nil
}

func createUnits(count uint64) []model.TicketUnit {
	res := make([]model.TicketUnit, 0, count)
	for range count {
		res = append(res, model.TicketUnit{Id: uuid.NewString(), Ticket: nil})
	}
	return res
}

func (repo *Memory) GetBooking(id string) *model.Booking {
	res, ok := repo.bookings[id]
	if !ok {
		return nil
	}
	return &res
}

func (repo *Memory) UpdateBooking(bk model.Booking) (*model.Booking, error) {
	repo.bookings[bk.Id] = bk
	return &bk, nil
}
