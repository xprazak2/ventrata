package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xprazak2/ventrata/internal/model"
	"github.com/xprazak2/ventrata/internal/repository"
	"github.com/xprazak2/ventrata/internal/repository/memory"
	"github.com/xprazak2/ventrata/internal/utils"
)

func setupProduct(
	t *testing.T, name string, price uint64, capacity uint64, currency string,
) (*model.Product, *model.Availability, *BookingService, repository.Repository) {
	repo := memory.NewMemoryRepository()

	bookingSvc := NewBookingService(repo)
	prod := repo.CreateEntry(name, capacity, price, currency)
	dateSvc := utils.NewDateProvider()
	availSvc := NewAvailabilityService(dateSvc, repo)
	avails, err := availSvc.GetAvailabilityRange(prod.Id, dateSvc.Today(), dateSvc.Today())
	assert.NoError(t, err)

	avail := avails[0]
	return prod, &avail, bookingSvc, repo
}

func TestCreateBooking(t *testing.T) {
	name := "foo"
	var capacity uint64 = 4
	var price uint64 = 200
	currency := "USD"
	var unitsCount uint64 = 3

	t.Run("should create a booking", func(t *testing.T) {
		prod, avail, bookingSvc, repo := setupProduct(t, name, price, capacity, currency)

		booking, err := bookingSvc.CreateBooking(prod.Id, avail.Id, unitsCount)
		assert.NoError(t, err)
		assert.Equal(t, prod.Id, booking.ProductId)
		assert.Equal(t, avail.Id, booking.AvailabilityId)
		assert.Equal(t, price, booking.PricePerUnit)
		assert.Len(t, booking.Units, int(unitsCount))
		assert.False(t, booking.IsConfirmed())
		assert.Nil(t, booking.Units[0].Ticket)
		assert.Equal(t, 600, int(booking.TotalPrice()))

		foundAvail, err := repo.GetAvailability(prod.Id, avail.Id)
		assert.NoError(t, err)
		assert.Equal(t, foundAvail.Vacancies, uint64(1))
	})

	t.Run("should not create a booking when there are insufficient vacancies", func(t *testing.T) {
		prod, avail, bookingSvc, repo := setupProduct(t, name, price, capacity, currency)

		_, err := bookingSvc.CreateBooking(prod.Id, avail.Id, 4)
		assert.NoError(t, err)

		foundAvail, err := repo.GetAvailability(prod.Id, avail.Id)
		assert.NoError(t, err)
		assert.Equal(t, foundAvail.Vacancies, uint64(0))
		assert.False(t, foundAvail.IsAvailable())
		assert.Equal(t, model.SOLD_OUT, foundAvail.Status())

		_, err = bookingSvc.CreateBooking(prod.Id, avail.Id, 2)
		assert.Error(t, err, "not enough vacancies")
	})
}

func TestConfirmBooking(t *testing.T) {
	name := "foo"
	var capacity uint64 = 4
	var price uint64 = 200
	currency := "USD"
	var unitsCount uint64 = 3

	t.Run("should confirm a booking and be idempotent", func(t *testing.T) {
		prod, avail, bookingSvc, _ := setupProduct(t, name, price, capacity, currency)
		booking, err := bookingSvc.CreateBooking(prod.Id, avail.Id, unitsCount)
		assert.False(t, booking.IsConfirmed())

		assert.NoError(t, err)

		confirmed, err := bookingSvc.ConfirmBooking(booking.Id)
		assert.NoError(t, err)
		assert.True(t, confirmed.IsConfirmed())

		confirmed2, err := bookingSvc.ConfirmBooking(confirmed.Id)
		assert.NoError(t, err)
		assert.True(t, confirmed2.IsConfirmed())
	})

	t.Run("should return nil for missing booking", func(t *testing.T) {
		_, _, bookingSvc, _ := setupProduct(t, name, price, capacity, currency)
		found, err := bookingSvc.ConfirmBooking("aaa")
		assert.Nil(t, found)
		assert.Nil(t, err)
	})
}
