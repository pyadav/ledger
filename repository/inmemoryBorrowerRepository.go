package repository

import "github.com/pyadav/navi/entity"

type InMemoryBorrowerRepository struct {
	borrowers map[string]*entity.Borrower
}

func NewInMemoryBorrowerRepository() *InMemoryBorrowerRepository {
	return &InMemoryBorrowerRepository{
		borrowers: make(map[string]*entity.Borrower),
	}
}

func (r *InMemoryBorrowerRepository) Save(borrower *entity.Borrower) {
	r.borrowers[borrower.Name] = borrower
}

func (r *InMemoryBorrowerRepository) FindByName(name string) *entity.Borrower {
	return r.borrowers[name]
}
