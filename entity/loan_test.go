package entity_test

import (
	"reflect"
	"testing"

	"github.com/pyadav/navi/entity"
	"github.com/stretchr/testify/assert"
)

func TestLoan_CalculateEmiAmount(t *testing.T) {
	tests := []struct {
		name              string
		loan              *entity.Loan
		expectedEmiAmount int
	}{
		{
			name: "Test with principal 10000, rate 4, and 5 years",
			loan: &entity.Loan{
				Principal:      5000,
				NoOfYears:      1,
				RateOfInterest: 6,
			},
			expectedEmiAmount: 442,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.loan.CalculateEmiAmount()
			assert.Equal(t, test.expectedEmiAmount, test.loan.EmiAmount)
		})
	}
}

func TestLoan_MakePayment(t *testing.T) {
	tests := []struct {
		name             string
		loan             *entity.Loan
		lumpSumAmount    int
		emiNo            int
		expectedPayments map[int]int
	}{
		{
			name: "Test with single lump sum payment after 5th EMI",
			loan: &entity.Loan{
				LumpSumPayments: make(map[int]int),
			},
			lumpSumAmount: 1000,
			emiNo:         5,
			expectedPayments: map[int]int{
				5: 1000,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.loan.MakePayment(test.lumpSumAmount, test.emiNo)

			assert.True(t, reflect.DeepEqual(test.expectedPayments, test.loan.LumpSumPayments))
		})
	}
}

func TestLoan_GetBalance(t *testing.T) {
	tests := []struct {
		name             string
		loan             *entity.Loan
		lumpSumPayments  map[int]int
		emiNo            int
		expectedAmount   int
		expectedEmisLeft int
	}{
		{
			name: "Test with lump sum payment after 3 EMIs",
			loan: &entity.Loan{
				Principal:      5000,
				NoOfYears:      1,
				RateOfInterest: 6,
			},
			lumpSumPayments:  map[int]int{5: 1000},
			emiNo:            3,
			expectedAmount:   1326,
			expectedEmisLeft: 9,
		},
		{
			name: "Test with lump sum payment after 6 EMIs",
			loan: &entity.Loan{
				Principal:      5000,
				NoOfYears:      1,
				RateOfInterest: 6,
			},
			lumpSumPayments:  map[int]int{5: 1000},
			emiNo:            6,
			expectedAmount:   3652,
			expectedEmisLeft: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.loan.LumpSumPayments = test.lumpSumPayments
			test.loan.CalculateEmiAmount()

			amountPaid, emisLeft := test.loan.GetBalance(test.emiNo)
			assert.Equal(t, test.expectedAmount, amountPaid)
			assert.Equal(t, test.expectedEmisLeft, emisLeft)
		})
	}
}
