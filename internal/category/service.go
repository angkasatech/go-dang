package category

import (
	"log"

	"go-dang/internal/errors"
)

// Service layer: bertanggung jawab untuk business logic
// Semua validation dan business rules ditangani di sini
// Semua error dari repository ditangani di sini
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetAll mengambil semua kategori
// Validation: tidak ada
// Error handling: dari repository
func (s *Service) GetAll() ([]Category, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		log.Printf("[SERVICE ERROR] GetAll failed: %v", err)
		return nil, errors.NewServiceError("gagal mengambil data kategori", err)
	}

	if categories == nil {
		categories = []Category{} // Return empty array instead of nil
	}

	return categories, nil
}

// GetByID mengambil kategori berdasarkan ID
// Validation: ID harus > 0
// Error handling: not found atau dari repository
func (s *Service) GetByID(id int) (*Category, error) {
	// Validation
	if id <= 0 {
		log.Printf("[SERVICE ERROR] Invalid ID: %d", id)
		return nil, errors.NewServiceError("ID kategori tidak valid", errors.ErrInvalidData)
	}

	category, err := s.repo.FindByID(id)
	if err != nil {
		log.Printf("[SERVICE ERROR] GetByID failed for ID %d: %v", id, err)
		return nil, errors.NewServiceError("gagal mengambil kategori", err)
	}

	return category, nil
}

// Create membuat kategori baru
// Validation: nama tidak boleh kosong, minimal 3 karakter
// Error handling: dari repository
func (s *Service) Create(req *CreateCategoryRequest) (*Category, error) {
	// Validation
	if err := req.Validate(); err != nil {
		log.Printf("[SERVICE ERROR] Create validation failed: %v", err)
		return nil, errors.NewServiceError("data kategori tidak valid: "+err.Error(), errors.ErrInvalidData)
	}

	category := &Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.repo.Create(category)
	if err != nil {
		log.Printf("[SERVICE ERROR] Create failed: %v", err)
		return nil, errors.NewServiceError("gagal menyimpan kategori", err)
	}

	return category, nil
}

// Update mengubah kategori
// Validation: ID harus > 0, nama tidak boleh kosong
// Error handling: not found atau dari repository
func (s *Service) Update(id int, req *UpdateCategoryRequest) (*Category, error) {
	// Validation
	if id <= 0 {
		log.Printf("[SERVICE ERROR] Invalid ID for update: %d", id)
		return nil, errors.NewServiceError("ID kategori tidak valid", errors.ErrInvalidData)
	}

	if err := req.Validate(); err != nil {
		log.Printf("[SERVICE ERROR] Update validation failed: %v", err)
		return nil, errors.NewServiceError("data kategori tidak valid: "+err.Error(), errors.ErrInvalidData)
	}

	// Cek apakah kategori exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		log.Printf("[SERVICE ERROR] Kategori not found for update: ID %d", id)
		return nil, errors.NewServiceError("kategori tidak ditemukan", err)
	}

	// Update data
	existing.Name = req.Name
	existing.Description = req.Description

	err = s.repo.Update(id, existing)
	if err != nil {
		log.Printf("[SERVICE ERROR] Update failed: %v", err)
		return nil, errors.NewServiceError("gagal mengubah kategori", err)
	}

	return existing, nil
}

// Delete menghapus kategori
// Validation: ID harus > 0
// Error handling: not found atau dari repository
func (s *Service) Delete(id int) error {
	// Validation
	if id <= 0 {
		log.Printf("[SERVICE ERROR] Invalid ID for delete: %d", id)
		return errors.NewServiceError("ID kategori tidak valid", errors.ErrInvalidData)
	}

	// Cek apakah kategori exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		log.Printf("[SERVICE ERROR] Kategori not found for delete: ID %d", id)
		return errors.NewServiceError("kategori tidak ditemukan", err)
	}

	err = s.repo.Delete(id)
	if err != nil {
		log.Printf("[SERVICE ERROR] Delete failed: %v", err)
		return errors.NewServiceError("gagal menghapus kategori", err)
	}

	return nil
}
