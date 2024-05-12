package memory

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/xprazak2/ventrata/internal/model"
	"github.com/xprazak2/ventrata/internal/utils"
	"github.com/xprazak2/ventrata/internal/verrors"
)

func (repo *Memory) UpdateAvailabilityVacancies(productId string, availId string, units uint64) error {
	entry := repo.getProductEntry(productId)
	if entry == nil {
		return verrors.NewNotFoundError("product", productId)
	}

	found, idx := findAvailability(entry.availability, availId)
	if found == nil {
		return verrors.NewNotFoundError("availability", availId)
	}

	entry.availability[idx] = model.Availability{
		Id:        found.Id,
		LocalDate: found.LocalDate,
		Vacancies: found.Vacancies - units,
	}

	return nil
}

func (repo *Memory) GetAvailability(productId string, availId string) (*model.Availability, error) {
	entry := repo.getProductEntry(productId)
	if entry == nil {
		return nil, verrors.NewNotFoundError("product", productId)
	}

	found, _ := findAvailability(entry.availability, availId)
	if found == nil {
		return nil, verrors.NewNotFoundError("availability", availId)
	}
	return found, nil
}

func findAvailability(avails []model.Availability, availId string) (*model.Availability, int) {
	for idx, item := range avails {
		if item.Id == availId {
			return &item, idx
		}
	}
	return nil, -1
}

func (repo *Memory) GetAvailabilityRange(
	productId string, startDate time.Time, endDate time.Time,
) ([]model.Availability, error) {
	entry := repo.getProductEntry(productId)
	if entry == nil {
		return nil, verrors.NewNotFoundError("product", productId)
	}

	added := []model.Availability{}
	count := len(entry.availability)

	first := entry.availability[0]

	endOffsetDays := offsetDays(endDate, first.LocalDate)

	// This method does too many things.
	// It should not need to check if sufficient range of availabilities is present
	// and it should not generate missing availability records.
	// Availability records for the whole year should be already created via a different mechanism
	// (regularly scheduled job once a day?)
	// and they should be already present so that
	// we can simply look up the existing records.
	// But the lifecycle of the product and is availabilities is not specified and in-memory is not a good storage anyway...
	if count < endOffsetDays {
		last := entry.availability[count-1]
		newDate := last.LocalDate.Add(24 * time.Hour)
		added = generateAvailabilityRange(
			newDate, endDate, entry.product.Capacity, entry.product.Price, entry.product.Currency,
		)
	}

	entry.availability = append(entry.availability, added...)

	startOffsetDays := offsetDays(startDate, first.LocalDate)

	return entry.availability[startOffsetDays : endOffsetDays+1], nil
}

func offsetDays(fst time.Time, snd time.Time) int {
	res := fst.Sub(snd)
	return int(math.Floor(res.Hours() / 24))
}

func generateAvailabilityYear(
	capacity uint64, price uint64, currency string,
) []model.Availability {
	dp := utils.NewDateProvider()
	sDate := dp.Today()
	eDate := sDate.Add(time.Hour * 24 * 365)
	return generateAvailabilityRange(sDate, eDate, capacity, price, currency)
}

func generateAvailabilityRange(
	startDate time.Time, endDate time.Time, cap uint64, price uint64, currency string,
) []model.Availability {
	newDate := startDate
	added := []model.Availability{}
	for {
		if newDate.After(endDate) {
			break
		}

		added = append(added, model.Availability{
			Id:        uuid.NewString(),
			LocalDate: newDate,
			Vacancies: cap,
			Price:     price,
			Currency:  currency,
		})

		newDate = newDate.Add(24 * time.Hour)
	}
	return added
}
