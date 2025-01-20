package tests

import (
	"testing"

	"github.com/Renan-Parise/finances/internal/api/statistics"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStatisticsUseCase struct {
	mock.Mock
}

func TestNewStatisticsHandler(t *testing.T) {
	router := gin.Default()
	group := router.Group("/api")

	mockUseCase := new(MockStatisticsUseCase)

	statistics.NewStatisticsHandler(group, mockUseCase)

	routes := []struct {
		method   string
		endpoint string
	}{
		{"GET", "/api/statistics/category-percentage"},
		{"GET", "/api/statistics/expenses-by-category"},
		{"GET", "/api/statistics/monthly-summary"},
		{"GET", "/api/statistics/highest-expenses"},
		{"GET", "/api/statistics/highest-incomes"},
		{"GET", "/api/statistics/spending-heatmap"},
		{"GET", "/api/statistics/general"},
	}

	for _, route := range routes {
		assert.True(t, routeExists(router, route.method, route.endpoint), "Route %s %s should exist", route.method, route.endpoint)
	}
}

func routeExists(router *gin.Engine, method, path string) bool {
	for _, route := range router.Routes() {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}

func (m *MockStatisticsUseCase) GetGeneralStatistics(userID int64) (*statistics.GeneralStatistics, error) {
	args := m.Called(userID)
	return args.Get(0).(*statistics.GeneralStatistics), args.Error(1)
}

func (m *MockStatisticsUseCase) GetExpensesByCategory(userID int64) ([]*statistics.ExpenseCategorySummary, error) {
	args := m.Called(userID)
	return args.Get(0).([]*statistics.ExpenseCategorySummary), args.Error(1)
}

func (m *MockStatisticsUseCase) GetCategoryPercentageChanges(userID int64) ([]*statistics.CategoryPercentageChange, error) {
	args := m.Called(userID)
	return args.Get(0).([]*statistics.CategoryPercentageChange), args.Error(1)
}

func (m *MockStatisticsUseCase) GetHighestExpenseMonth(userID int64) (*statistics.MonthlyAmount, error) {
	args := m.Called(userID)
	return args.Get(0).(*statistics.MonthlyAmount), args.Error(1)
}

func (m *MockStatisticsUseCase) GetHighestIncomeMonth(userID int64) (*statistics.MonthlyAmount, error) {
	args := m.Called(userID)
	return args.Get(0).(*statistics.MonthlyAmount), args.Error(1)
}

func (m *MockStatisticsUseCase) GetMonthlyExpensesSummary(userID int64) ([]*statistics.MonthlyAmount, error) {
	args := m.Called(userID)
	return args.Get(0).([]*statistics.MonthlyAmount), args.Error(1)
}

func (m *MockStatisticsUseCase) GetSpendingHeatmap(userID int64) (map[string]float64, error) {
	args := m.Called(userID)
	return args.Get(0).(map[string]float64), args.Error(1)
}
