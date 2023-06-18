package repository

import "github.com/pyadav/navi/entity"

type InMemoryLoanRepository struct {
	lenders   map[string]*entity.Lender
	borrowers map[string]*entity.Borrower
}

func NewInMemoryLoanRepository() *InMemoryLoanRepository {
	return &InMemoryLoanRepository{
		lenders:   make(map[string]*entity.Lender),
		borrowers: make(map[string]*entity.Borrower),
	}
}

func (r *InMemoryLoanRepository) SaveLoan(lender *entity.Lender, borrower *entity.Borrower, loan *entity.Loan) {
	lender.Loans = append(lender.Loans, loan)
	borrower.Loans = append(borrower.Loans, loan)
}

func (r *InMemoryLoanRepository) FindByLenderAndBorrower(lender *entity.Lender, borrower *entity.Borrower) *entity.Loan {
	for _, loan := range lender.Loans {
		if loan.Borrower == borrower {
			return loan
		}
	}
	return nil
}
