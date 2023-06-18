package repository

import "github.com/pyadav/navi/entity"

type BorrowerRepository interface {
	Save(borrower *entity.Borrower)
	FindByName(name string) *entity.Borrower
}
