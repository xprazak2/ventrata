package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/xprazak2/ventrata/internal/model"
	"github.com/xprazak2/ventrata/internal/repository"
)

type BookingService struct {
	repo repository.Repository
}

func NewBookingService(repo repository.Repository) *BookingService {
	return &BookingService{repo}
}

func (svc *BookingService) CreateBooking(productId string, availId string, units uint64) (*model.Booking, error) {
	if units == 0 {
		return nil, fmt.Errorf("at least one unit required for a booking to be created")
	}

	svc.repo.Lock()
	defer svc.repo.Unlock()

	avail, err := svc.repo.GetAvailability(productId, availId)
	if err != nil {
		return nil, err
	}

	if avail.Vacancies < units {
		return nil, fmt.Errorf("not enough vacancies")
	}

	err = svc.repo.UpdateAvailabilityVacancies(productId, availId, units)
	if err != nil {
		return nil, err
	}

	booking, err := svc.repo.CreateBooking(productId, availId, units, avail.Price, avail.Currency)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (svc *BookingService) GetBooking(id string) *model.Booking {
	return svc.repo.GetBooking(id)
}

func (svc *BookingService) ConfirmBooking(id string) (*model.Booking, error) {
	svc.repo.Lock()
	defer svc.repo.Unlock()

	found := svc.repo.GetBooking(id)
	if found == nil {
		return nil, nil
	}

	if found.IsConfirmed() {
		return found, nil
	}

	found.Units = generateTickets(found.Units)
	found.Status = model.CONFIRMED

	return svc.repo.UpdateBooking(*found)
}

func generateTickets(units []model.TicketUnit) []model.TicketUnit {
	res := make([]model.TicketUnit, 0, len(units))
	for _, item := range units {
		id := uuid.NewString()
		res = append(res, model.TicketUnit{Id: item.Id, Ticket: &id})
	}
	return res
}
