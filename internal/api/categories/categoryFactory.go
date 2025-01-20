package categories

import (
	"time"
)

func NewCategory(userID int64, name string) *Category {
	now := time.Now()
	return &Category{
		UserID:    userID,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
