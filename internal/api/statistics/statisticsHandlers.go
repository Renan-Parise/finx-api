package statistics

import (
	"net/http"

	"github.com/Renan-Parise/finances/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	statisticsUseCase StatisticsUseCase
}

func NewStatisticsHandler(router *gin.RouterGroup, su StatisticsUseCase) {
	handler := &StatisticsHandler{
		statisticsUseCase: su,
	}

	statistics := router.Group("/statistics")
	statistics.Use(middlewares.JWTAuthMiddleware())
	statistics.Use(middlewares.RedisCacheMiddleware)
	{
		statistics.GET("/category-percentage", handler.GetCategoryPercentageChanges)
		statistics.GET("/expenses-by-category", handler.GetExpensesByCategory)
		statistics.GET("/monthly-summary", handler.GetMonthlyExpensesSummary)
		statistics.GET("/highest-expenses", handler.GetHighestExpenseMonth)
		statistics.GET("/highest-incomes", handler.GetHighestIncomeMonth)
		statistics.GET("/spending-heatmap", handler.GetSpendingHeatmap)
		statistics.GET("/general", handler.GetGeneralStatistics)
	}
}

func (h *StatisticsHandler) GetGeneralStatistics(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	stats, err := h.statisticsUseCase.GetGeneralStatistics(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *StatisticsHandler) GetHighestExpenseMonth(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	monthData, err := h.statisticsUseCase.GetHighestExpenseMonth(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if monthData == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No expense data available"})
		return
	}

	c.JSON(http.StatusOK, monthData)
}

func (h *StatisticsHandler) GetHighestIncomeMonth(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	monthData, err := h.statisticsUseCase.GetHighestIncomeMonth(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if monthData == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No income data available"})
		return
	}

	c.JSON(http.StatusOK, monthData)
}

func (h *StatisticsHandler) GetCategoryPercentageChanges(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	changes, err := h.statisticsUseCase.GetCategoryPercentageChanges(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, changes)
}

func (h *StatisticsHandler) GetSpendingHeatmap(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	heatmap, err := h.statisticsUseCase.GetSpendingHeatmap(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, heatmap)
}

func (h *StatisticsHandler) GetMonthlyExpensesSummary(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	summary, err := h.statisticsUseCase.GetMonthlyExpensesSummary(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (h *StatisticsHandler) GetExpensesByCategory(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	summary, err := h.statisticsUseCase.GetExpensesByCategory(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
