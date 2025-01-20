package categories

import (
	"database/sql"

	"github.com/Renan-Parise/finances/internal/errors"
)

type CategoryRepository interface {
	Create(category *Category) error
	GetAll(userID int64) ([]*Category, error)
	Delete(userID int64, id int) error
	ExistsByName(userID int64, name string) (bool, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *Category) error {
	query := `INSERT INTO categories (userId, name, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errors.NewQueryError("error preparing query: " + err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Exec(category.UserID, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return errors.NewQueryError("error executing query: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.NewQueryError("error getting last insert ID: " + err.Error())
	}

	category.ID = int(id)
	return nil
}

func (r *categoryRepository) GetAll(userID int64) ([]*Category, error) {
	query := `SELECT id, userId, name, createdAt, updatedAt FROM categories WHERE userId = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("error executing query: " + err.Error())
	}
	defer rows.Close()

	var categories []*Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, errors.NewQueryError("error scanning row: " + err.Error())
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *categoryRepository) Delete(userID int64, id int) error {
	query := `DELETE FROM categories WHERE id = ? AND userId = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errors.NewQueryError("error preparing query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userID)
	return err
}

func (r *categoryRepository) ExistsByName(userID int64, name string) (bool, error) {
	query := `SELECT COUNT(*) FROM categories WHERE userId = ? AND LOWER(name) = LOWER(?)`
	var count int
	err := r.db.QueryRow(query, userID, name).Scan(&count)
	if err != nil {
		return false, errors.NewQueryError("error executing query: " + err.Error())
	}
	return count > 0, nil
}
