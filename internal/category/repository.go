package category

import (
	"database/sql"
	"log"

	"go-dang/internal/errors"
)

// Repository layer: bertanggung jawab untuk akses data dari database
// Semua error dari database ditangani di sini
type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// FindAll mengambil semua kategori dari database
// Error yang kembalikan adalah dari database operation
func (r *Repository) FindAll() ([]Category, error) {
	rows, err := r.db.Query("SELECT id, name, description, created_at, updated_at FROM categories")
	if err != nil {
		log.Printf("[REPOSITORY ERROR] FindAll query failed: %v", err)
		return nil, errors.NewRepositoryError("gagal mengambil data kategori", err)
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		scanErr := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
		if scanErr != nil {
			log.Printf("[REPOSITORY ERROR] Scan failed: %v", scanErr)
			return nil, errors.NewRepositoryError("gagal membaca data kategori", scanErr)
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[REPOSITORY ERROR] Rows iteration failed: %v", err)
		return nil, errors.NewRepositoryError("gagal iterasi data kategori", err)
	}

	return categories, nil
}

// FindByID mengambil kategori berdasarkan ID
// Error: not found atau database error
func (r *Repository) FindByID(id int) (*Category, error) {
	var c Category
	err := r.db.QueryRow(
		"SELECT id, name, description, created_at, updated_at FROM categories WHERE id=$1",
		id,
	).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)

	if err == sql.ErrNoRows {
		log.Printf("[REPOSITORY ERROR] Category ID %d not found", id)
		return nil, errors.NewRepositoryError("kategori tidak ditemukan", errors.ErrNotFound)
	}

	if err != nil {
		log.Printf("[REPOSITORY ERROR] FindByID failed for ID %d: %v", id, err)
		return nil, errors.NewRepositoryError("gagal mengambil kategori", err)
	}

	return &c, nil
}

// Create menambah kategori baru ke database
// Error: validation atau database error
func (r *Repository) Create(c *Category) error {
	err := r.db.QueryRow(
		"INSERT INTO categories(name, description, created_at, updated_at) VALUES($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at",
		c.Name, c.Description,
	).Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)

	if err != nil {
		log.Printf("[REPOSITORY ERROR] Create failed: %v", err)
		return errors.NewRepositoryError("gagal menyimpan kategori ke database", err)
	}

	return nil
}

// Update mengubah kategori di database
// Error: not found atau database error
func (r *Repository) Update(id int, c *Category) error {
	result, err := r.db.Exec(
		"UPDATE categories SET name=$1, description=$2, updated_at=NOW() WHERE id=$3",
		c.Name, c.Description, id,
	)

	if err != nil {
		log.Printf("[REPOSITORY ERROR] Update failed for ID %d: %v", id, err)
		return errors.NewRepositoryError("gagal mengubah kategori di database", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[REPOSITORY ERROR] RowsAffected failed: %v", err)
		return errors.NewRepositoryError("gagal mengecek hasil update", err)
	}

	if rowsAffected == 0 {
		log.Printf("[REPOSITORY ERROR] No rows updated for ID %d", id)
		return errors.NewRepositoryError("kategori tidak ditemukan", errors.ErrNotFound)
	}

	return nil
}

// Delete menghapus kategori dari database
// Error: not found atau database error
func (r *Repository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM categories WHERE id=$1", id)

	if err != nil {
		log.Printf("[REPOSITORY ERROR] Delete failed for ID %d: %v", id, err)
		return errors.NewRepositoryError("gagal menghapus kategori dari database", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[REPOSITORY ERROR] RowsAffected failed: %v", err)
		return errors.NewRepositoryError("gagal mengecek hasil delete", err)
	}

	if rowsAffected == 0 {
		log.Printf("[REPOSITORY ERROR] No rows deleted for ID %d", id)
		return errors.NewRepositoryError("kategori tidak ditemukan", errors.ErrNotFound)
	}

	return nil
}
