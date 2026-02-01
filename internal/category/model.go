package category

import (
	"errors"
	"time"
)

// Category adalah domain model untuk data kategori
// Gunakan struct ini untuk merepresentasikan data di seluruh layer
type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateCategoryRequest adalah request body untuk create category
type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateCategoryRequest adalah request body untuk update category
type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate untuk memastikan data update category valid
func (u *UpdateCategoryRequest) Validate() error {
	if u.Name == "" {
		return errors.New("nama kategori tidak boleh kosong")
	}
	if len(u.Name) < 3 {
		return errors.New("nama kategori minimal 3 karakter")
	}
	return nil
}

// CategoryResponse adalah response body untuk category
type CategoryResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate untuk memastikan data category valid
func (c *CreateCategoryRequest) Validate() error {
	if c.Name == "" {
		return errors.New("nama kategori tidak boleh kosong")
	}
	if len(c.Name) < 3 {
		return errors.New("nama kategori minimal 3 karakter")
	}
	return nil
}
