package main

import (
	"github.com/Renan-Parise/finances/internal/api/categories"
	"github.com/Renan-Parise/finances/internal/api/statistics"
	"github.com/Renan-Parise/finances/internal/api/transactions"
	"github.com/Renan-Parise/finances/internal/container"
	"github.com/Renan-Parise/finances/internal/db"
	"github.com/Renan-Parise/finances/internal/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	db.RunMigrations()
	redis.GetRedis()

	container := container.NewContainer()

	router := gin.Default()

	api := router.Group("/api")

	transactions.NewTransactionHandler(api, container.TransactionUseCase)
	statistics.NewStatisticsHandler(api, container.StatisticsUseCase)
	categories.NewCategoryHandler(api, container.CategoryUseCase)

	router.Run("0.0.0.0:8180")
}
