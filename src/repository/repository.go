package repository

import (
	"gorestapi/src/domain"
	"sync"
)

type ItemRepository interface {
	GetAll() []domain.Item
	Add(item domain.Item)
}

type ItemInMemoryRepository struct {
	Items   []domain.Item
	rwMutex sync.Mutex
}

func NewItemInMemoryRepository() *ItemInMemoryRepository {
	return &ItemInMemoryRepository{
		Items: []domain.Item{},
	}
}

func (r *ItemInMemoryRepository) Add(item domain.Item) {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	r.Items = append(r.Items, item)
}

func (r *ItemInMemoryRepository) GetAll() []domain.Item {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	return r.Items
}
