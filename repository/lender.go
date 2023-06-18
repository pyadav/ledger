package repository

import "github.com/pyadav/navi/entity"

type LenderRepository interface {
	Save(lender *entity.Lender)
	FindByName(name string) *entity.Lender
}
