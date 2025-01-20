package tests

import (
	"testing"

	"github.com/Renan-Parise/finances/internal/api/categories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryUseCase struct {
	mock.Mock
}

func TestNewCategoryHandler(t *testing.T) {
	router := gin.Default()
	group := router.Group("/api")

	mockUseCase := new(MockCategoryUseCase)
	categories.NewCategoryHandler(group, mockUseCase)

	routes := router.Routes()

	expectedRoutes := []struct {
		Method string
		Path   string
	}{
		{"POST", "/api/categories/"},
		{"GET", "/api/categories/"},
		{"DELETE", "/api/categories/:id"},
		{"POST", "/api/categories/default"},
	}

	for _, expected := range expectedRoutes {
		found := false
		for _, route := range routes {
			if route.Method == expected.Method && route.Path == expected.Path {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected route [%s] %s not found", expected.Method, expected.Path)
	}
}

func (m *MockCategoryUseCase) CreateCategory(id int64, name string) error {
	args := m.Called(id, name)
	return args.Error(0)
}

func (m *MockCategoryUseCase) CreateDefaultCategories(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCategoryUseCase) DeleteCategory(id int64, userID int) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *MockCategoryUseCase) GetCategories(userID int64) ([]*categories.Category, error) {
	args := m.Called(userID)
	return args.Get(0).([]*categories.Category), args.Error(1)
}
