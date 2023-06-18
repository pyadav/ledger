package entity

import "math"

type Loan struct {
	Lender          *Lender
	Borrower        *Borrower
	Principal       int
	NoOfYears       int
	RateOfInterest  int
	EmiAmount       int
	LumpSumPayments map[int]int
}

func (l *Loan) CalculateEmiAmount() {
	totalAmount := l.Principal + (l.Principal*l.RateOfInterest*l.NoOfYears)/100
	noOfEmis := int(float64(l.NoOfYears * 12))
	l.EmiAmount = int(math.Ceil(float64(totalAmount) / float64(noOfEmis)))
}

func (l *Loan) MakePayment(lumpSumAmount, emiNo int) {
	l.LumpSumPayments[emiNo] = lumpSumAmount
}

func (l *Loan) GetBalance(emiNo int) (int, int) {
	amountPaid := 0
	for i := 1; i <= emiNo; i++ {
		amountPaid += l.EmiAmount
		if lumpSum, ok := l.LumpSumPayments[i]; ok {
			amountPaid += lumpSum
		}
	}

	emisRemaining := int(math.Ceil(float64(l.Principal+(l.Principal*l.RateOfInterest*l.NoOfYears)/100-amountPaid) / float64(l.EmiAmount)))
	return amountPaid, emisRemaining
}
