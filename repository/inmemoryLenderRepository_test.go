package repository_test

import (
	"testing"

	"github.com/pyadav/navi/entity"
	"github.com/pyadav/navi/repository"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryLenderRepository_Save(t *testing.T) {
	repo := repository.NewInMemoryLenderRepository()

	lender := &entity.Lender{Name: "John"}
	repo.Save(lender)

	savedLender := repo.FindByName("John")

	assert.NotNil(t, savedLender)
	assert.Equal(t, savedLender, lender)
}

func TestInMemoryLenderRepository_FindByName(t *testing.T) {
	repo := repository.NewInMemoryLenderRepository()

	lender1 := &entity.Lender{Name: "John"}
	lender2 := &entity.Lender{Name: "Jane"}

	repo.Save(lender1)
	repo.Save(lender2)

	foundLender1 := repo.FindByName("John")
	foundLender2 := repo.FindByName("Jane")
	notFoundLender := repo.FindByName("Alice")

	assert.Equal(t, foundLender1, lender1)
	assert.Equal(t, foundLender2, lender2)
	assert.Nil(t, notFoundLender)
}
