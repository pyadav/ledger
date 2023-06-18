package repository

import "github.com/pyadav/navi/entity"

type InMemoryLenderRepository struct {
	lenders map[string]*entity.Lender
}

func NewInMemoryLenderRepository() *InMemoryLenderRepository {
	return &InMemoryLenderRepository{
		lenders: make(map[string]*entity.Lender),
	}
}

func (r *InMemoryLenderRepository) Save(lender *entity.Lender) {
	r.lenders[lender.Name] = lender
}

func (r *InMemoryLenderRepository) FindByName(name string) *entity.Lender {
	return r.lenders[name]
}
