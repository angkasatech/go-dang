package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, _ := h.service.GetAll()
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, _ := h.service.GetByID(id)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c Category
	json.NewDecoder(r.Body).Decode(&c)
	h.service.Create(&c)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var c Category
	json.NewDecoder(r.Body).Decode(&c)
	h.service.Update(id, &c)
	json.NewEncoder(w).Encode(c)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	h.service.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
