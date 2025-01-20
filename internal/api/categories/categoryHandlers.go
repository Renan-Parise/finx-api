package categories

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Renan-Parise/finances/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUseCase CategoryUseCase
}

func NewCategoryHandler(router *gin.RouterGroup, cu CategoryUseCase) {
	handler := &CategoryHandler{
		categoryUseCase: cu,
	}

	categories := router.Group("/categories")
	categories.Use(middlewares.JWTAuthMiddleware())
	{
		categories.POST("/", handler.CreateCategory)
		categories.GET("/", handler.GetCategories)
		categories.DELETE("/:id", handler.DeleteCategory)
	}

	categories.POST("/default", handler.CreateDefaultCategories)
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.categoryUseCase.CreateCategory(userID.(int64), input.Name)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	categories, err := h.categoryUseCase.GetCategories(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = h.categoryUseCase.DeleteCategory(userID.(int64), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *CategoryHandler) CreateDefaultCategories(c *gin.Context) {
	var input struct {
		UserID int64 `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	err := h.categoryUseCase.CreateDefaultCategories(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Default categories created successfully"})
}
