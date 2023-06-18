package service

import (
	"github.com/pyadav/navi/entity"
	"github.com/pyadav/navi/repository"
)

type LoanService struct {
	lenderRepo   repository.LenderRepository
	borrowerRepo repository.BorrowerRepository
	loanRepo     repository.LoanRepository
}

func NewLoanService(
	lenderRepo repository.LenderRepository,
	borrowerRepo repository.BorrowerRepository,
	loanRepo repository.LoanRepository,
) *LoanService {
	return &LoanService{
		lenderRepo:   lenderRepo,
		borrowerRepo: borrowerRepo,
		loanRepo:     loanRepo,
	}
}

func (s *LoanService) CreateLoan(lenderName, borrowerName string, principal, noOfYears, rateOfInterest int) {
	lender := s.lenderRepo.FindByName(lenderName)
	if lender == nil {
		lender = &entity.Lender{Name: lenderName}
		s.lenderRepo.Save(lender)
	}

	borrower := s.borrowerRepo.FindByName(borrowerName)
	if borrower == nil {
		borrower = &entity.Borrower{Name: borrowerName}
		s.borrowerRepo.Save(borrower)
	}

	loan := &entity.Loan{
		Lender:          lender,
		Borrower:        borrower,
		Principal:       principal,
		NoOfYears:       noOfYears,
		RateOfInterest:  rateOfInterest,
		LumpSumPayments: make(map[int]int),
	}

	loan.CalculateEmiAmount()
	s.loanRepo.SaveLoan(lender, borrower, loan)
}

func (s *LoanService) MakePayment(lenderName, borrowerName string, lumpSumAmount, emiNo int) {
	lender := s.lenderRepo.FindByName(lenderName)
	borrower := s.borrowerRepo.FindByName(borrowerName)
	if lender != nil && borrower != nil {
		loan := s.loanRepo.FindByLenderAndBorrower(lender, borrower)
		if loan != nil {
			loan.MakePayment(lumpSumAmount, emiNo)
		}
	}
}

func (s *LoanService) GetBalance(lenderName, borrowerName string, emiNo int) (int, int) {
	lender := s.lenderRepo.FindByName(lenderName)
	borrower := s.borrowerRepo.FindByName(borrowerName)
	if lender != nil && borrower != nil {
		loan := s.loanRepo.FindByLenderAndBorrower(lender, borrower)
		if loan != nil {
			return loan.GetBalance(emiNo)
		}
	}
	return 0, 0
}
