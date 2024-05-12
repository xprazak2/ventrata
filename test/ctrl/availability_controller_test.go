package ctrl

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xprazak2/ventrata/internal/controller"
	"github.com/xprazak2/ventrata/internal/utils"
	"github.com/xprazak2/ventrata/internal/view"
)

func TestGetAvailability(t *testing.T) {
	t.Run("should get availability", func(t *testing.T) {
		id := "foo"
		resp := []view.Availability{}
		params := controller.DateRangeParams{
			ProductId: id,
			LocalDate: &utils.LocalDate{Time: time.Now().Add(time.Hour * 24)},
		}
		err := Request("POST", "/availability", &params, &resp, false)
		if err != nil {
			t.Fail()
		}

		assert.Len(t, resp, 1)
		assert.Equal(t, uint64(5), resp[0].Vacancies)
	})

	t.Run("should get priced availability", func(t *testing.T) {
		id := "foo"
		resp := []view.PricedAvailability{}
		params := controller.DateRangeParams{
			ProductId: id,
			LocalDate: &utils.LocalDate{Time: time.Now().Add(time.Hour * 24)},
		}
		err := Request("POST", "/availability", &params, &resp, true)
		if err != nil {
			t.Fail()
		}

		assert.Len(t, resp, 1)
		assert.Equal(t, "EUR", resp[0].Currency)
	})
}
