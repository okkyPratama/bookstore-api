package repository

import (
	"database/sql"

	"github.com/okkyPratama/bookstore-api/structs"
)

type CategoryRepository struct {
	DB *sql.DB
}

func (r *CategoryRepository) GetAllCategories() ([]structs.Category, error) {
	query := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var c structs.Category
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, &c.ModifiedAt, &c.ModifiedBy)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) CreateCategory(category *structs.Category) error {
	query := `INSERT INTO categories (name, created_by) VALUES ($1, $2) RETURNING id`
	return r.DB.QueryRow(query, category.Name, category.CreatedBy).Scan(&category.ID)
}

func (r *CategoryRepository) GetCategoryByID(id int) (*structs.Category, error) {
	category := &structs.Category{}
	query := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(category *structs.Category) error {
	query := `UPDATE categories SET name = $1, modified_at = CURRENT_TIMESTAMP, modified_by = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, category.Name, category.ModifiedBy, category.ID)
	return err
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
