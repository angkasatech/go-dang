package category

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go-dang/internal/errors"

	"github.com/go-chi/chi/v5"
)

// Handler layer: bertanggung jawab untuk menerima request dan mengirim response
// Semua HTTP-specific logic ditangani di sini
type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Helper function untuk send JSON response
func (h *Handler) sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Helper function untuk send error response
func (h *Handler) sendError(w http.ResponseWriter, statusCode int, message string, err error) {
	log.Printf("[HANDLER ERROR] %s: %v", message, err)

	errorResp := map[string]interface{}{
		"error":   message,
		"details": err.Error(),
	}

	h.sendJSON(w, statusCode, errorResp)
}

// GetAll mengambil semua kategori
// GET /categories
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("[HANDLER] GetAll called")

	data, err := h.service.GetAll()
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, "Gagal mengambil data kategori", err)
		return
	}

	resp := map[string]interface{}{
		"status": "success",
		"data":   data,
		"count":  len(data),
	}

	h.sendJSON(w, http.StatusOK, resp)
}

// GetByID mengambil kategori berdasarkan ID
// GET /categories/:id
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	log.Printf("[HANDLER] GetByID called with ID: %s", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, "ID kategori tidak valid", err)
		return
	}

	data, err := h.service.GetByID(id)
	if err != nil {
		// Check jika error adalah "not found"
		if _, ok := err.(*errors.AppError); ok {
			appErr := err.(*errors.AppError)
			if appErr.Err == errors.ErrNotFound {
				h.sendError(w, http.StatusNotFound, "Kategori tidak ditemukan", err)
				return
			}
		}
		h.sendError(w, http.StatusInternalServerError, "Gagal mengambil kategori", err)
		return
	}

	resp := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	h.sendJSON(w, http.StatusOK, resp)
}

// Create membuat kategori baru
// POST /categories
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("[HANDLER] Create called")

	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, http.StatusBadRequest, "Format request tidak valid", err)
		return
	}

	data, err := h.service.Create(&req)
	if err != nil {
		// Check jika error adalah validation error
		if _, ok := err.(*errors.AppError); ok {
			appErr := err.(*errors.AppError)
			if appErr.Err == errors.ErrInvalidData {
				h.sendError(w, http.StatusBadRequest, appErr.Message, err)
				return
			}
		}
		h.sendError(w, http.StatusInternalServerError, "Gagal membuat kategori", err)
		return
	}

	resp := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	h.sendJSON(w, http.StatusCreated, resp)
}

// Update mengubah kategori
// PUT /categories/:id
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	log.Printf("[HANDLER] Update called with ID: %s", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, "ID kategori tidak valid", err)
		return
	}

	var req UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, http.StatusBadRequest, "Format request tidak valid", err)
		return
	}

	data, err := h.service.Update(id, &req)
	if err != nil {
		// Check jika error adalah "not found" atau validation error
		if _, ok := err.(*errors.AppError); ok {
			appErr := err.(*errors.AppError)
			if appErr.Err == errors.ErrNotFound {
				h.sendError(w, http.StatusNotFound, "Kategori tidak ditemukan", err)
				return
			}
			if appErr.Err == errors.ErrInvalidData {
				h.sendError(w, http.StatusBadRequest, appErr.Message, err)
				return
			}
		}
		h.sendError(w, http.StatusInternalServerError, "Gagal mengubah kategori", err)
		return
	}

	resp := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	h.sendJSON(w, http.StatusOK, resp)
}

// Delete menghapus kategori
// DELETE /categories/:id
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	log.Printf("[HANDLER] Delete called with ID: %s", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, "ID kategori tidak valid", err)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		// Check jika error adalah "not found"
		if _, ok := err.(*errors.AppError); ok {
			appErr := err.(*errors.AppError)
			if appErr.Err == errors.ErrNotFound {
				h.sendError(w, http.StatusNotFound, "Kategori tidak ditemukan", err)
				return
			}
		}
		h.sendError(w, http.StatusInternalServerError, "Gagal menghapus kategori", err)
		return
	}

	resp := map[string]interface{}{
		"status":  "success",
		"message": "Kategori berhasil dihapus",
	}

	h.sendJSON(w, http.StatusOK, resp)
}
