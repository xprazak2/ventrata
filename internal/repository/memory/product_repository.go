package memory

import "github.com/xprazak2/ventrata/internal/model"

func (repo *Memory) GetProducts() []model.Product {
	res := make([]model.Product, 0, len(repo.entries))

	for _, entry := range repo.entries {
		res = append(res, entry.product)
	}

	return res
}

func (repo *Memory) GetProduct(id string) *model.Product {
	found := repo.getProductEntry(id)
	if found == nil {
		return nil
	}
	return &found.product
}

func (repo *Memory) getProductEntry(id string) *Entry {
	found, ok := repo.entries[id]
	if !ok {
		return nil
	}
	return &found
}
