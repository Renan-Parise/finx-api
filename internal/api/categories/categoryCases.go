package categories

import (
	"github.com/Renan-Parise/finances/internal/errors"
)

type CategoryUseCase interface {
	GetCategories(userID int64) ([]*Category, error)
	CreateCategory(userID int64, name string) error
	CreateDefaultCategories(userID int64) error
	DeleteCategory(userID int64, id int) error
}

type categoryUseCase struct {
	categoryRepo CategoryRepository
}

func NewCategoryUseCase(cr CategoryRepository) CategoryUseCase {
	return &categoryUseCase{categoryRepo: cr}
}

func (uc *categoryUseCase) CreateCategory(userID int64, name string) error {
	exists, err := uc.categoryRepo.ExistsByName(userID, name)
	if err != nil {
		return errors.NewServiceError("error checking if category name exists: " + err.Error())
	}

	if exists {
		return errors.NewValidationError(name, "the given category name already exists: "+name)
	}

	category := NewCategory(userID, name)
	return uc.categoryRepo.Create(category)
}

func (uc *categoryUseCase) CreateDefaultCategories(userID int64) error {
	defaultCategories := []string{"Food", "Entertainment", "Transport", "Shopping", "Salary", "Travel"}

	for _, categoryName := range defaultCategories {
		category := NewCategory(userID, categoryName)
		if err := uc.categoryRepo.Create(category); err != nil {
			return errors.NewServiceError("error creating category " + categoryName + ": " + err.Error())
		}
	}
	return nil
}

func (uc *categoryUseCase) GetCategories(userID int64) ([]*Category, error) {
	return uc.categoryRepo.GetAll(userID)
}

func (uc *categoryUseCase) DeleteCategory(userID int64, id int) error {
	return uc.categoryRepo.Delete(userID, id)
}
