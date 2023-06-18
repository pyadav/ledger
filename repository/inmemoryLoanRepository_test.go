package repository_test

import (
	"testing"

	"github.com/pyadav/navi/entity"
	"github.com/pyadav/navi/repository"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryLoanRepository_SaveLoan(t *testing.T) {
	repo := repository.NewInMemoryLoanRepository()

	lender := &entity.Lender{Name: "John"}
	borrower := &entity.Borrower{Name: "Alice"}

	loan := &entity.Loan{
		Lender:          lender,
		Borrower:        borrower,
		Principal:       1000,
		NoOfYears:       5,
		RateOfInterest:  8,
		EmiAmount:       0,
		LumpSumPayments: map[int]int{},
	}

	repo.SaveLoan(lender, borrower, loan)

	assert.Equal(t, len(lender.Loans), 1)
	assert.Equal(t, len(borrower.Loans), 1)
	assert.Equal(t, borrower.Loans[0], loan)
	assert.Equal(t, borrower.Loans[0], loan)
}

func TestInMemoryLoanRepository_FindByLenderAndBorrower(t *testing.T) {
	repo := repository.NewInMemoryLoanRepository()

	lender := &entity.Lender{Name: "John"}
	borrower := &entity.Borrower{Name: "Alice"}
	loan := &entity.Loan{
		Lender:          lender,
		Borrower:        borrower,
		Principal:       1000,
		NoOfYears:       5,
		RateOfInterest:  8,
		EmiAmount:       0,
		LumpSumPayments: map[int]int{},
	}

	repo.SaveLoan(lender, borrower, loan)

	foundLoan := repo.FindByLenderAndBorrower(lender, borrower)
	if foundLoan != loan {
		t.Errorf("FindByLenderAndBorrower failed: incorrect loan returned")
	}

	otherBorrower := &entity.Borrower{Name: "Bob"}
	notFoundLoan := repo.FindByLenderAndBorrower(lender, otherBorrower)

	assert.Nil(t, notFoundLoan)
}
