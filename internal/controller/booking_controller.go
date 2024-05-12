package controller

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xprazak2/ventrata/internal/service"
	"github.com/xprazak2/ventrata/internal/view"
)

type BookingParams struct {
	ProductId      string `json:"product_id"`
	AvailabilityId string `json:"availability_id"`
	Units          uint64 `json:"units"`
}

func CreateBooking(svc *service.BookingService) func(w http.ResponseWriter, t *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPriced := IsPriced(r.Context())

		var params BookingParams
		err := parseParams(&params, w, r)
		if err != nil {
			return
		}

		booking, err := svc.CreateBooking(params.ProductId, params.AvailabilityId, params.Units)
		if err != nil {
			ResponseError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		ResponseJSON(w, view.BookingView(*booking, isPriced), http.StatusCreated)
	}
}

func GetBooking(svc *service.BookingService) func(w http.ResponseWriter, t *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPriced := IsPriced(r.Context())
		bookingId := chi.URLParam(r, "id")

		booking := svc.GetBooking(bookingId)
		if booking == nil {
			ResponseError(w, fmt.Sprintf("Booking with id '%s' not found", bookingId), http.StatusNotFound)
			return
		}

		ResponseJSON(w, view.BookingView(*booking, isPriced), http.StatusOK)
	}
}

func ConfirmBooking(svc *service.BookingService) func(w http.ResponseWriter, t *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPriced := IsPriced(r.Context())
		bookingId := chi.URLParam(r, "id")

		booking, err := svc.ConfirmBooking(bookingId)
		if err != nil {
			ResponseError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if booking == nil {
			ResponseError(w, fmt.Sprintf("Booking with id '%s' not found", bookingId), http.StatusNotFound)
			return
		}

		ResponseJSON(w, view.BookingView(*booking, isPriced), http.StatusCreated)
	}
}
