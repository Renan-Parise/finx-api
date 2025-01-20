package statistics

type GeneralStatistics struct {
	TotalIncome      float64 `json:"totalIncome"`
	TotalExpenses    float64 `json:"totalExpenses"`
	Balance          float64 `json:"balance"`
	MostUsedCategory string  `json:"mostUsedCategory"`
}

type MonthlyAmount struct {
	Year  int     `json:"year"`
	Month int     `json:"month"`
	Total float64 `json:"total"`
}

type CategoryPercentageChange struct {
	CategoryName     string  `json:"categoryName"`
	PreviousValue    float64 `json:"previousValue"`
	CurrentValue     float64 `json:"currentValue"`
	PercentageChange float64 `json:"percentageChange"`
	Increase         bool    `json:"increase"`
}

type ExpenseCategorySummary struct {
	CategoryName string  `json:"categoryName"`
	TotalAmount  float64 `json:"totalAmount"`
	Percentage   float64 `json:"percentage"`
}
