package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/xprazak2/ventrata/internal/service"
	"github.com/xprazak2/ventrata/internal/utils"
	"github.com/xprazak2/ventrata/internal/verrors"
	"github.com/xprazak2/ventrata/internal/view"
)

type DateRangeParams struct {
	ProductId      string           `json:"product_id"`
	LocalDate      *utils.LocalDate `json:"local_date"`
	LocalDateStart *utils.LocalDate `json:"local_date_start"`
	LocalDateEnd   *utils.LocalDate `json:"local_date_end"`
}

func GetAvailability(svc *service.AvailabilityService) func(w http.ResponseWriter, t *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPriced := IsPriced(r.Context())

		var rangeParams DateRangeParams
		err := parseParams(&rangeParams, w, r)
		if err != nil {
			return
		}

		startDate, endDate, err := getRangeDates(rangeParams)
		if err != nil {
			http.Error(w, "invalid request parameters", http.StatusBadRequest)
			return
		}

		avail, err := svc.GetAvailabilityRange(rangeParams.ProductId, *startDate, *endDate)

		if err != nil {
			status := http.StatusBadRequest
			var notFound *verrors.NotFoundError
			if errors.As(err, &notFound) {
				status = http.StatusNotFound
			}

			ResponseError(w, err.Error(), status)
			return
		}
		ResponseJSON(w, view.AvailabilityView(avail, isPriced), http.StatusOK)
	}
}

func getRangeDates(params DateRangeParams) (*time.Time, *time.Time, error) {
	if params.LocalDate != nil {
		return &params.LocalDate.Time, &params.LocalDate.Time, nil
	} else if params.LocalDateStart != nil && params.LocalDateEnd != nil {
		return &params.LocalDateStart.Time, &params.LocalDateEnd.Time, nil
	} else {
		return nil, nil, fmt.Errorf("invalid request parameters")
	}
}
