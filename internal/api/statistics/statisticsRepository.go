package statistics

import (
	"database/sql"

	"github.com/Renan-Parise/finances/internal/errors"
)

type StatisticsRepository interface {
	GetCategoryMonthlyTotals(userID int64, month, year int) (map[string]float64, error)
	GetExpensesByCategory(userID int64) ([]*ExpenseCategorySummary, error)
	GetMonthlyExpensesSummary(userID int64) ([]*MonthlyAmount, error)
	GetMonthlyExpenses(userID int64) ([]*MonthlyAmount, error)
	GetMonthlyIncome(userID int64) ([]*MonthlyAmount, error)
	GetSpendingHeatmap(userID int64) (map[string]float64, error)
	GetMostUsedCategory(userID int64) (string, error)
	GetTotalExpenses(userID int64) (float64, error)
	GetTotalIncome(userID int64) (float64, error)
}

type statisticsRepository struct {
	db *sql.DB
}

func NewStatisticsRepository(db *sql.DB) StatisticsRepository {
	return &statisticsRepository{db: db}
}

func (r *statisticsRepository) GetTotalIncome(userID int64) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE userId = ? AND amount > 0`
	var totalIncome float64
	err := r.db.QueryRow(query, userID).Scan(&totalIncome)
	return totalIncome, err
}

func (r *statisticsRepository) GetTotalExpenses(userID int64) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE userId = ? AND amount < 0`
	var totalExpenses float64
	err := r.db.QueryRow(query, userID).Scan(&totalExpenses)
	return totalExpenses, err
}

func (r *statisticsRepository) GetMostUsedCategory(userID int64) (string, error) {
	query := `
		SELECT c.name, COUNT(*) AS usage_count
		FROM transactions t
		JOIN categories c ON t.category = c.id
		WHERE t.userId = ?
		GROUP BY t.category
		ORDER BY usage_count DESC
		LIMIT 1
	`
	var categoryName string
	var usageCount int

	err := r.db.QueryRow(query, userID).Scan(&categoryName, &usageCount)
	return categoryName, err
}

func (r *statisticsRepository) GetMonthlyExpenses(userID int64) ([]*MonthlyAmount, error) {
	query := `
		SELECT YEAR(createdAt) as year, MONTH(createdAt) as month, ABS(SUM(amount)) as total
		FROM transactions
		WHERE userId = ? AND amount < 0
		GROUP BY YEAR(createdAt), MONTH(createdAt)
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("Failed to get monthly expenses: " + err.Error())
	}
	defer rows.Close()

	var results []*MonthlyAmount
	for rows.Next() {
		var ma MonthlyAmount
		err := rows.Scan(&ma.Year, &ma.Month, &ma.Total)
		if err != nil {
			return nil, errors.NewQueryError("Failed to scan monthly expenses: " + err.Error())
		}
		results = append(results, &ma)
	}
	return results, nil
}

func (r *statisticsRepository) GetMonthlyIncome(userID int64) ([]*MonthlyAmount, error) {
	query := `
		SELECT YEAR(createdAt) as year, MONTH(createdAt) as month, SUM(amount) as total
		FROM transactions
		WHERE userId = ? AND amount > 0
		GROUP BY YEAR(createdAt), MONTH(createdAt)
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("Failed to get monthly income: " + err.Error())
	}
	defer rows.Close()

	var results []*MonthlyAmount
	for rows.Next() {
		var ma MonthlyAmount
		err := rows.Scan(&ma.Year, &ma.Month, &ma.Total)
		if err != nil {
			return nil, errors.NewQueryError("Failed to scan monthly income: " + err.Error())
		}
		results = append(results, &ma)
	}
	return results, nil
}

func (r *statisticsRepository) GetCategoryMonthlyTotals(userID int64, month, year int) (map[string]float64, error) {
	query := `
		SELECT c.name, SUM(t.amount) as total
		FROM transactions t
		JOIN categories c ON t.category = c.id
		WHERE t.userId = ? AND MONTH(t.createdAt) = ? AND YEAR(t.createdAt) = ?
		GROUP BY c.name
	`
	rows, err := r.db.Query(query, userID, month, year)
	if err != nil {
		return nil, errors.NewQueryError("Failed to get category monthly totals: " + err.Error())
	}
	defer rows.Close()

	totals := make(map[string]float64)
	for rows.Next() {
		var categoryName string
		var total float64
		err := rows.Scan(&categoryName, &total)
		if err != nil {
			return nil, errors.NewQueryError("Failed to scan category monthly totals: " + err.Error())
		}
		totals[categoryName] = total
	}
	return totals, nil
}

func (r *statisticsRepository) GetSpendingHeatmap(userID int64) (map[string]float64, error) {
	query := `
		SELECT DATE(createdAt) as day, ABS(SUM(amount)) as total
		FROM transactions
		WHERE userId = ? AND amount < 0 AND createdAt >= DATE_SUB(CURRENT_DATE, INTERVAL 11 MONTH)
		GROUP BY DATE(createdAt)
		ORDER BY day ASC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("Failed to get spending heatmap: " + err.Error())
	}
	defer rows.Close()

	heatmap := make(map[string]float64)
	for rows.Next() {
		var day string
		var total float64
		err := rows.Scan(&day, &total)
		if err != nil {
			return nil, errors.NewQueryError("Failed to scan heatmap data: " + err.Error())
		}
		heatmap[day] = total
	}
	return heatmap, nil
}

func (r *statisticsRepository) GetMonthlyExpensesSummary(userID int64) ([]*MonthlyAmount, error) {
	query := `
		SELECT YEAR(createdAt) as year, MONTH(createdAt) as month, ABS(SUM(amount)) as total
		FROM transactions
		WHERE userId = ? AND amount < 0
		GROUP BY YEAR(createdAt), MONTH(createdAt)
		ORDER BY year DESC, month DESC
		LIMIT 12
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("Failed to get monthly expenses summary: " + err.Error())
	}
	defer rows.Close()

	var results []*MonthlyAmount
	for rows.Next() {
		var ma MonthlyAmount
		err := rows.Scan(&ma.Year, &ma.Month, &ma.Total)
		if err != nil {
			return nil, errors.NewQueryError("Failed to scan monthly expenses summary: " + err.Error())
		}
		results = append(results, &ma)
	}
	return results, nil
}

func (r *statisticsRepository) GetExpensesByCategory(userID int64) ([]*ExpenseCategorySummary, error) {
	query := `
		SELECT c.name, ABS(SUM(t.amount)) as total
		FROM transactions t
		JOIN categories c ON t.category = c.id
		WHERE t.userId = ? AND t.amount < 0
		GROUP BY c.name
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("Failed to get expenses by category: " + err.Error())
	}
	defer rows.Close()

	var results []*ExpenseCategorySummary
	var totalSum float64

	for rows.Next() {
		var category string
		var total float64
		err := rows.Scan(&category, &total)
		if err != nil {
			return nil, errors.NewQueryError("Failed to scan category expenses: " + err.Error())
		}
		totalSum += total
		results = append(results, &ExpenseCategorySummary{
			CategoryName: category,
			TotalAmount:  total,
		})
	}

	for _, result := range results {
		result.Percentage = (result.TotalAmount / totalSum) * 100
	}

	return results, nil
}
