package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/xprazak2/ventrata/internal/model"
)

type Entry struct {
	product      model.Product
	availability []model.Availability
}

type Memory struct {
	lock     sync.Mutex
	entries  map[string]Entry
	bookings map[string]model.Booking
}

func NewMemoryRepository() *Memory {
	return &Memory{entries: map[string]Entry{}, bookings: map[string]model.Booking{}}
}

func (repo *Memory) Lock() {
	repo.lock.Lock()
}

func (repo *Memory) Unlock() {
	repo.lock.Unlock()
}

func (repo *Memory) CreateEntry(name string, capacity uint64, price uint64, currency string) *model.Product {
	id := uuid.NewString()
	return repo.createEntry(id, name, capacity, price, currency)
}

func (repo *Memory) createEntry(
	id string, name string, capacity uint64, price uint64, currency string,
) *model.Product {
	product := model.Product{
		Id:       id,
		Name:     name,
		Capacity: capacity,
		Price:    price,
		Currency: currency,
	}

	repo.entries[id] = Entry{
		product:      product,
		availability: generateAvailabilityYear(capacity, price, currency),
	}

	return &product
}

func (repo *Memory) SeedData() {
	repo.createEntry("foo", "City Underground", 5, 1000, "EUR")
	repo.createEntry("bar", "Bus Tour", 3, 1500, "USD")
}
