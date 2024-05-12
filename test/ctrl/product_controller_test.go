package ctrl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xprazak2/ventrata/internal/view"
)

func TestGetProduct(t *testing.T) {
	t.Run("should get product", func(t *testing.T) {
		id := "foo"
		resp := view.Product{}
		err := Request("GET", "/products/"+id, nil, &resp, false)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, id, resp.Id)
	})

	t.Run("should get priced product", func(t *testing.T) {
		id := "foo"
		resp := view.PricedProduct{}
		err := Request("GET", "/products/"+id, nil, &resp, true)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, id, resp.Id)
		assert.Equal(t, uint64(1000), resp.Price)
		assert.Equal(t, "EUR", resp.Currency)
	})
}

func TestGetProducts(t *testing.T) {
	t.Run("should get products", func(t *testing.T) {
		resp := []view.Product{}
		err := Request("GET", "/products", nil, &resp, false)
		if err != nil {
			t.Fail()
		}

		assert.Len(t, resp, 2)
	})

	t.Run("should get priced products", func(t *testing.T) {
		resp := []view.PricedProduct{}
		err := Request("GET", "/products", nil, &resp, true)
		if err != nil {
			t.Fail()
		}

		assert.Len(t, resp, 2)
		for _, prod := range resp {
			if prod.Id == "foo" {
				assert.Equal(t, uint64(1000), prod.Price)
				assert.Equal(t, "EUR", prod.Currency)
			}
		}
	})
}
