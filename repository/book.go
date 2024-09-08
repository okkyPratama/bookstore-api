package repository

import (
	"database/sql"

	"github.com/okkyPratama/bookstore-api/structs"
)

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) GetAllBooks() ([]structs.Book, error) {
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var b structs.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *BookRepository) CreateBook(book *structs.Book) error {
	query := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	return r.DB.QueryRow(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedBy).Scan(&book.ID)
}

func (r *BookRepository) GetBookByID(id int) (*structs.Book, error) {
	book := &structs.Book{}
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by 
              FROM books WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) UpdateBook(book *structs.Book) error {
	query := `UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, modified_at = CURRENT_TIMESTAMP, modified_by = $9 WHERE id = $10`
	_, err := r.DB.Exec(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.ModifiedBy, book.ID)
	return err
}

func (r *BookRepository) DeleteBook(id int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *BookRepository) GetBooksByCategory(categoryID int) ([]structs.Book, error) {
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by 
              FROM books WHERE category_id = $1`
	rows, err := r.DB.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var b structs.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
