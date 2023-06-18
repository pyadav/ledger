package repository

import "github.com/pyadav/navi/entity"

type LoanRepository interface {
	SaveLoan(lender *entity.Lender, borrower *entity.Borrower, loan *entity.Loan)
	FindByLenderAndBorrower(lender *entity.Lender, borrower *entity.Borrower) *entity.Loan
}
