package tests

import (
	"testing"

	"github.com/Renan-Parise/finances/internal/api/transactions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUseCase struct {
	mock.Mock
}

func TestNewTransactionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	group := router.Group("/api")

	mockUseCase := new(MockTransactionUseCase)
	transactions.NewTransactionHandler(group, mockUseCase)

	routes := router.Routes()

	expectedRoutes := []struct {
		method string
		path   string
	}{
		{"POST", "/api/transactions/filter"},
		{"DELETE", "/api/transactions/:id"},
		{"PUT", "/api/transactions/:id"},
		{"POST", "/api/transactions/"},
		{"GET", "/api/transactions/"},
	}

	for _, expected := range expectedRoutes {
		found := false
		for _, route := range routes {
			if route.Method == expected.method && route.Path == expected.path {
				found = true
				break
			}
		}
		assert.True(t, found, "Route %s %s not registered", expected.method, expected.path)
	}
}

func (m *MockTransactionUseCase) DeleteTransaction(id1 int64, id2 int64) error {
	args := m.Called(id1, id2)
	return args.Error(0)
}

func (m *MockTransactionUseCase) FilterTransactions(id int64, filter *transactions.Filter) ([]*transactions.Transaction, error) {
	args := m.Called(id, filter)
	return args.Get(0).([]*transactions.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) CreateTransaction(id int64, name string, quantity int, price float64) error {
	args := m.Called(id, name, quantity, price)
	return args.Error(0)
}

func (m *MockTransactionUseCase) GetTransactions(id int64) ([]*transactions.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).([]*transactions.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) UpdateTransaction(transaction *transactions.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
