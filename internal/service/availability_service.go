package service

import (
	"fmt"
	"time"

	"github.com/xprazak2/ventrata/internal/model"
	"github.com/xprazak2/ventrata/internal/repository"
	"github.com/xprazak2/ventrata/internal/utils"
)

type AvailabilityService struct {
	dateProvider utils.DateProvider
	repo         repository.Repository
}

func NewAvailabilityService(dateProvider utils.DateProvider, repo repository.Repository) *AvailabilityService {
	return &AvailabilityService{dateProvider, repo}
}

func (svc *AvailabilityService) GetAvailabilityRange(
	productId string, startDate time.Time, endDate time.Time,
) ([]model.Availability, error) {
	if endDate.Before(startDate) {
		return nil, fmt.Errorf("invalid date range")
	}

	today := svc.dateProvider.Today()
	maxDate := today.Add(time.Hour * 24 * 365)

	if startDate.After(maxDate) {
		return []model.Availability{}, nil
	}

	if endDate.Before(today) {
		return []model.Availability{}, nil
	}

	if startDate.Before(today) {
		startDate = today
	}

	if endDate.After(maxDate) {
		endDate = maxDate
	}

	svc.repo.Lock()
	defer svc.repo.Unlock()

	return svc.repo.GetAvailabilityRange(productId, startDate, endDate)
}
