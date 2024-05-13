package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xprazak2/ventrata/internal/repository/memory"
	"github.com/xprazak2/ventrata/internal/utils"
)

func TestGetAvailabilityRange(t *testing.T) {
	repo := memory.NewMemoryRepository()

	tdp := utils.NewDateProvider()

	svc := NewAvailabilityService(tdp, repo)
	product := repo.CreateEntry("foo", 5, 500, "EUR")

	t.Run("should get a single point", func(t *testing.T) {
		day := tdp.Today().Add(24 * time.Hour)
		avl, err := svc.GetAvailabilityRange(product.Id, day, day)
		assert.NoError(t, err)
		assert.Len(t, avl, 1)
		assert.Equal(t, avl[0].LocalDate, day)
	})

	t.Run("should get a range", func(t *testing.T) {
		startDay := tdp.Today().Add(24 * time.Hour)
		endDay := tdp.Today().Add(3 * 24 * time.Hour)
		avl, err := svc.GetAvailabilityRange(product.Id, startDay, endDay)
		assert.NoError(t, err)
		assert.Len(t, avl, 3)
		assert.Equal(t, avl[0].LocalDate, startDay)
		assert.Equal(t, avl[2].LocalDate, endDay)
	})

	t.Run("should return error when start is after end", func(t *testing.T) {
		endDay := tdp.Today().Add(24 * time.Hour)
		startDay := tdp.Today().Add(3 * 24 * time.Hour)
		_, err := svc.GetAvailabilityRange(product.Id, startDay, endDay)
		assert.Error(t, err, "invalid date range")
	})

	t.Run("should return nothing when end is in the past", func(t *testing.T) {
		endDay := tdp.Today().Add(-24 * time.Hour)
		startDay := tdp.Today().Add(-3 * 24 * time.Hour)
		avl, err := svc.GetAvailabilityRange(product.Id, startDay, endDay)
		assert.NoError(t, err)
		assert.Len(t, avl, 0)
	})

	t.Run("should return nothing when start is in more than a year in the future", func(t *testing.T) {
		startDay := tdp.Today().Add(400 * 24 * time.Hour)
		endDay := tdp.Today().Add(402 * 24 * time.Hour)
		avl, err := svc.GetAvailabilityRange(product.Id, startDay, endDay)
		assert.NoError(t, err)
		assert.Len(t, avl, 0)
	})

	t.Run("should return only items starting today when start is in the past", func(t *testing.T) {
		startDay := tdp.Today().Add(-1 * 24 * time.Hour)
		endDay := tdp.Today().Add(1 * 24 * time.Hour)
		avl, err := svc.GetAvailabilityRange(product.Id, startDay, endDay)
		assert.NoError(t, err)
		assert.Len(t, avl, 2)
		assert.Equal(t, avl[0].LocalDate, tdp.Today())
	})

	t.Run("should return only items that are at most a year in the future", func(t *testing.T) {
		startDay := tdp.Today().Add(364 * 24 * time.Hour)
		endDay := tdp.Today().Add(370 * 24 * time.Hour)
		avl, err := svc.GetAvailabilityRange(product.Id, startDay, endDay)
		assert.NoError(t, err)
		assert.Len(t, avl, 2)
	})
}
