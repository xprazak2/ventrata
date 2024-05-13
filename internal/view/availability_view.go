package view

import (
	"github.com/xprazak2/ventrata/internal/model"
	"github.com/xprazak2/ventrata/internal/utils"
)

type Availability struct {
	Id        string                   `json:"id"`
	LocalDate utils.LocalDate          `json:"local_date"`
	Status    model.AvailabilityStatus `json:"status"`
	Vacancies uint64                   `json:"vacancies"`
	Available bool                     `json:"available"`
}

type PricedAvailability struct {
	Availability
	Price    uint64 `json:"price"`
	Currency string `json:"currency"`
}

func AvailabilityView(items []model.Availability, isPriced bool) []interface{} {
	res := make([]interface{}, len(items))

	for idx, item := range items {

		avl := Availability{
			Id:        item.Id,
			LocalDate: utils.LocalDate{Time: item.LocalDate},
			Status:    item.Status(),
			Vacancies: item.Vacancies,
			Available: item.IsAvailable(),
		}

		if isPriced {
			res[idx] = PricedAvailability{
				Availability: avl,
				Price:        item.Price,
				Currency:     item.Currency,
			}
		} else {
			res[idx] = avl
		}
	}

	return res
}
