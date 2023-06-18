package repository_test

import (
	"testing"

	"github.com/pyadav/navi/entity"
	"github.com/pyadav/navi/repository"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryBorrowerRepository_Save(t *testing.T) {
	repo := repository.NewInMemoryBorrowerRepository()

	borrower := &entity.Borrower{Name: "Alice"}
	repo.Save(borrower)

	savedBorrower := repo.FindByName("Alice")

	assert.NotNil(t, savedBorrower)
	assert.Equal(t, savedBorrower, borrower)
}

func TestInMemoryBorrowerRepository_FindByName(t *testing.T) {
	repo := repository.NewInMemoryBorrowerRepository()

	borrower1 := &entity.Borrower{Name: "Alice"}
	borrower2 := &entity.Borrower{Name: "Bob"}

	repo.Save(borrower1)
	repo.Save(borrower2)

	foundBorrower1 := repo.FindByName("Alice")
	foundBorrower2 := repo.FindByName("Bob")
	notFoundBorrower := repo.FindByName("Eve")

	assert.Equal(t, foundBorrower1, borrower1)
	assert.Equal(t, foundBorrower2, borrower2)
	assert.Nil(t, notFoundBorrower)
}
