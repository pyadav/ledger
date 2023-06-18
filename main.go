package main

import (
	"fmt"
	"os"

	"github.com/pyadav/navi/processor"
	"github.com/pyadav/navi/repository"
	"github.com/pyadav/navi/service"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		return
	}

	filePath := cliArgs[0]

	lenderRepo := repository.NewInMemoryLenderRepository()
	borroeweRepo := repository.NewInMemoryBorrowerRepository()
	loanRepo := repository.NewInMemoryLoanRepository()

	loanService := service.NewLoanService(lenderRepo, borroeweRepo, loanRepo)

	err := processor.ProcessCommands(filePath, loanService)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
}
