package service_test

import (
	"testing"

	"github.com/pyadav/navi/entity"
	"github.com/pyadav/navi/service"
	"github.com/stretchr/testify/assert"
)

type mockLenderRepository struct {
	lenders map[string]*entity.Lender
}

func (r *mockLenderRepository) Save(lender *entity.Lender) {
	r.lenders[lender.Name] = lender
}

func (r *mockLenderRepository) FindByName(name string) *entity.Lender {
	return r.lenders[name]
}

type mockBorrowerRepository struct {
	borrowers map[string]*entity.Borrower
}

func (r *mockBorrowerRepository) Save(borrower *entity.Borrower) {
	r.borrowers[borrower.Name] = borrower
}

func (r *mockBorrowerRepository) FindByName(name string) *entity.Borrower {
	return r.borrowers[name]
}

type mockLoanRepository struct {
	loans map[string]*entity.Loan
}

func (r *mockLoanRepository) SaveLoan(lender *entity.Lender, borrower *entity.Borrower, loan *entity.Loan) {
	key := lender.Name + borrower.Name
	r.loans[key] = loan
}

func (r *mockLoanRepository) FindByLenderAndBorrower(lender *entity.Lender, borrower *entity.Borrower) *entity.Loan {
	key := lender.Name + borrower.Name
	return r.loans[key]
}

func TestLoanService_CreateLoan(t *testing.T) {
	lenderRepo := &mockLenderRepository{lenders: make(map[string]*entity.Lender)}
	borrowerRepo := &mockBorrowerRepository{borrowers: make(map[string]*entity.Borrower)}
	loanRepo := &mockLoanRepository{loans: make(map[string]*entity.Loan)}

	loanService := service.NewLoanService(lenderRepo, borrowerRepo, loanRepo)

	loanService.CreateLoan("John", "Alice", 10000, 5, 10)
	lender := lenderRepo.FindByName("John")
	borrower := borrowerRepo.FindByName("Alice")

	assert.NotNil(t, lender)
	assert.NotNil(t, borrower)

	loan := loanRepo.FindByLenderAndBorrower(lender, borrower)

	assert.NotNil(t, loan)
	assert.Equal(t, loan.Principal, 10000)
	assert.Equal(t, loan.NoOfYears, 5)
	assert.Equal(t, loan.RateOfInterest, 10)
}

func TestLoanService_MakePayment(t *testing.T) {
	lenderRepo := &mockLenderRepository{lenders: make(map[string]*entity.Lender)}
	borrowerRepo := &mockBorrowerRepository{borrowers: make(map[string]*entity.Borrower)}
	loanRepo := &mockLoanRepository{loans: make(map[string]*entity.Loan)}

	loanService := service.NewLoanService(lenderRepo, borrowerRepo, loanRepo)

	lender := &entity.Lender{Name: "John"}
	borrower := &entity.Borrower{Name: "Alice"}
	loan := &entity.Loan{
		Lender:          lender,
		Borrower:        borrower,
		Principal:       15000,
		NoOfYears:       2,
		RateOfInterest:  9,
		LumpSumPayments: make(map[int]int),
	}
	loan.CalculateEmiAmount()

	lenderRepo.Save(lender)
	borrowerRepo.Save(borrower)
	loanRepo.SaveLoan(lender, borrower, loan)

	loanService.MakePayment("John", "Alice", 7000, 12)

	amountPaid, emisRemaining := loan.GetBalance(12)

	assert.Equal(t, amountPaid, 15856)
	assert.Equal(t, emisRemaining, 3)
}

func TestLoanService_GetBalance(t *testing.T) {
	lenderRepo := &mockLenderRepository{lenders: make(map[string]*entity.Lender)}
	borrowerRepo := &mockBorrowerRepository{borrowers: make(map[string]*entity.Borrower)}
	loanRepo := &mockLoanRepository{loans: make(map[string]*entity.Loan)}

	loanService := service.NewLoanService(lenderRepo, borrowerRepo, loanRepo)

	lender := &entity.Lender{Name: "John"}
	borrower := &entity.Borrower{Name: "Alice"}
	loan := &entity.Loan{
		Lender:          lender,
		Borrower:        borrower,
		Principal:       10000,
		NoOfYears:       5,
		RateOfInterest:  10,
		LumpSumPayments: make(map[int]int),
	}
	loan.CalculateEmiAmount()

	lenderRepo.Save(lender)
	borrowerRepo.Save(borrower)
	loanRepo.SaveLoan(lender, borrower, loan)

	amountPaid, emisRemaining := loanService.GetBalance("John", "Alice", 3)

	assert.Equal(t, amountPaid, 750)
	assert.Equal(t, emisRemaining, 57)
}
