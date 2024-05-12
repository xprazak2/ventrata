package repository

import (
	"time"

	"github.com/xprazak2/ventrata/internal/model"
)

type Repository interface {
	GetProducts() []model.Product
	GetProduct(id string) *model.Product
	GetAvailabilityRange(productId string, startDate time.Time, endDate time.Time) ([]model.Availability, error)
	GetAvailability(productId string, availId string) (*model.Availability, error)
	CreateBooking(productId string, availId string, units uint64, price uint64, currency string) (*model.Booking, error)
	GetBooking(id string) *model.Booking
	UpdateAvailabilityVacancies(productId string, availId string, units uint64) error
	UpdateBooking(bk model.Booking) (*model.Booking, error)
	Lock()
	Unlock()
}
