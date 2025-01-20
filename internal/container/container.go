package container

import (
	"github.com/Renan-Parise/finances/internal/api/categories"
	"github.com/Renan-Parise/finances/internal/api/statistics"
	"github.com/Renan-Parise/finances/internal/api/transactions"
	"github.com/Renan-Parise/finances/internal/db"
)

type Container struct {
	CategoryRepository categories.CategoryRepository
	CategoryUseCase    categories.CategoryUseCase

	TransactionRepository transactions.TransactionRepositories
	TransactionUseCase    transactions.TransactionUseCase

	StatisticsRepository statistics.StatisticsRepository
	StatisticsUseCase    statistics.StatisticsUseCase
}

func NewContainer() *Container {
	database := db.GetDB()

	categoryRepo := categories.NewCategoryRepository(database)
	statisticsRepo := statistics.NewStatisticsRepository(database)
	transactionRepo := transactions.NewTransactionRepositories(database)

	categoryUseCase := categories.NewCategoryUseCase(categoryRepo)
	statisticsUseCase := statistics.NewStatisticsUseCase(statisticsRepo)
	transactionUseCase := transactions.NewTransactionUseCase(transactionRepo)

	return &Container{
		TransactionUseCase:    transactionUseCase,
		TransactionRepository: transactionRepo,

		CategoryUseCase:    categoryUseCase,
		CategoryRepository: categoryRepo,

		StatisticsUseCase:    statisticsUseCase,
		StatisticsRepository: statisticsRepo,
	}
}
