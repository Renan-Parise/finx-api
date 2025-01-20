package transactions

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Description string    `json:"description"`
	Category    int       `json:"category"`
	Amount      float64   `json:"amount"`
}

type Filter struct {
	Category int    `json:"category"`
	Search   string `json:"search"`
	Order    string `json:"order"`
	Field    string `json:"field"`
	From     string `json:"from"`
	File     string `json:"file"`
	To       string `json:"to"`
}
