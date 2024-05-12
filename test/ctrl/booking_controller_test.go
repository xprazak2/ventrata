package ctrl

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xprazak2/ventrata/internal/controller"
	"github.com/xprazak2/ventrata/internal/utils"
	"github.com/xprazak2/ventrata/internal/view"
)

func TestBookingController(t *testing.T) {
	id := "foo"
	availResp := []view.Availability{}
	availParams := controller.DateRangeParams{
		ProductId: id,
		LocalDate: &utils.LocalDate{Time: time.Now().Add(time.Hour * 24)},
	}

	err := Request("POST", "/availability", &availParams, &availResp, false)
	if err != nil {
		t.Fail()
	}
	assert.Len(t, availResp, 1)

	avail := availResp[0]

	bookingParams := controller.BookingParams{ProductId: id, AvailabilityId: avail.Id, Units: 5}
	bookingResp := view.Booking{}
	err = Request("POST", "/bookings", &bookingParams, &bookingResp, false)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, id, bookingResp.ProductId)
	assert.Equal(t, avail.Id, bookingResp.AvailabilityId)
	assert.Equal(t, 0, bookingResp.Status)
	assert.Len(t, bookingResp.Units, 5)
	assert.Nil(t, bookingResp.Units[0].Ticket)

	err = Request("POST", "/availability", &availParams, &availResp, false)
	if err != nil {
		t.Fail()
	}

	avail = availResp[0]
	assert.Equal(t, uint64(0), avail.Vacancies)

	errResp := controller.ErrorResp{}
	err = Request("POST", "/bookings", &bookingParams, &errResp, false)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "not enough vacancies", errResp.Error)
	confirmResp := view.PricedBooking{}
	err = Request("POST", "/bookings/"+bookingResp.Id+"/confirm", nil, &confirmResp, true)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 1, confirmResp.Status)
	assert.Equal(t, uint64(5000), confirmResp.Price)
	assert.Equal(t, uint64(1000), confirmResp.Units[0].Price)
	assert.NotNil(t, confirmResp.Units[0].Ticket)
}
