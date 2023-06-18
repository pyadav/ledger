package processor

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pyadav/navi/service"
)

const (
	LOAN    string = "LOAN"
	PAYMENT string = "PAYMENT"
	BALANCE string = "BALANCE"
)

func ProcessCommands(name string, svc *service.LoanService) error {
	content, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		serviceType := tokens[0]

		switch serviceType {
		case LOAN:
			bankName := tokens[1]
			borrowerName := tokens[2]
			principal, _ := strconv.Atoi(tokens[3])
			noOfYears, _ := strconv.Atoi(tokens[4])
			rateOfInterest, _ := strconv.Atoi(tokens[5])
			svc.CreateLoan(bankName, borrowerName, principal, noOfYears, rateOfInterest)

		case PAYMENT:
			bankName := tokens[1]
			borrowerName := tokens[2]
			lumpSumAmount, _ := strconv.Atoi(tokens[3])
			emiNo, _ := strconv.Atoi(tokens[4])
			svc.MakePayment(bankName, borrowerName, lumpSumAmount, emiNo)

		case BALANCE:
			bankName := tokens[1]
			borrowerName := tokens[2]
			emiNo, _ := strconv.Atoi(tokens[3])
			amountPaid, emisRemaining := svc.GetBalance(bankName, borrowerName, emiNo)
			fmt.Printf("%s %s %d %d\n", bankName, borrowerName, amountPaid, emisRemaining)
		}
	}

	return nil
}
