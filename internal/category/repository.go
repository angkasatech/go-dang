package category

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		rows.Scan(&c.ID, &c.Name, &c.Description)
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *Repository) FindByID(id int) (*Category, error) {
	var c Category
	err := r.db.QueryRow(
		"SELECT id, name, description FROM categories WHERE id=$1", id,
	).Scan(&c.ID, &c.Name, &c.Description)

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *Repository) Create(c *Category) error {
	return r.db.QueryRow(
		"INSERT INTO categories(name, description) VALUES($1,$2) RETURNING id",
		c.Name, c.Description,
	).Scan(&c.ID)
}

func (r *Repository) Update(id int, c *Category) error {
	_, err := r.db.Exec(
		"UPDATE categories SET name=$1, description=$2 WHERE id=$3",
		c.Name, c.Description, id,
	)
	return err
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id=$1", id)
	return err
}
