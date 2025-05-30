// internal/api/handlers/farmer_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samuel-prates/farm-project/backend/internal/models"
	"github.com/samuel-prates/farm-project/backend/pkg/logger"
)

type FarmerHandler struct {
	service FarmerServiceInterface
}

func NewFarmerHandler(service FarmerServiceInterface) *FarmerHandler {
	return &FarmerHandler{service: service}
}

func (h *FarmerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var farmer models.Farmer
	if err := json.NewDecoder(r.Body).Decode(&farmer); err != nil {
		logger.Warn("Erro ao decodificar JSON: %v", err)
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := farmer.Validate(); err != nil {
		logger.Warn("Erro de validação ao criar fazendeiro: %v", err)
		http.Error(w, "Erro de validação: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdFarmer, err := h.service.Create(&farmer)
	if err != nil {
		logger.Error("Erro ao criar fazendeiro: %v", err)
		http.Error(w, "Erro ao criar fazendeiro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFarmer)
}

func (h *FarmerHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		logger.Warn("ID inválido ao atualizar fazendeiro: %v", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var farmer models.Farmer
	if err := json.NewDecoder(r.Body).Decode(&farmer); err != nil {
		logger.Warn("Erro ao decodificar JSON na atualização de fazendeiro: %v", err)
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	farmer.ID = uint(id)
	if err := farmer.Validate(); err != nil {
		logger.Warn("Erro de validação ao atualizar fazendeiro: %v", err)
		http.Error(w, "Erro de validação: "+err.Error(), http.StatusBadRequest)
		return
	}

	updatedFarmer, err := h.service.Update(&farmer)
	if err != nil {
		logger.Error("Erro ao atualizar fazendeiro: %v", err)
		http.Error(w, "Erro ao atualizar fazendeiro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedFarmer)
}

func (h *FarmerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		logger.Warn("ID inválido ao excluir fazendeiro: %v", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		logger.Error("Erro ao excluir fazendeiro: %v", err)
		http.Error(w, "Erro ao excluir fazendeiro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *FarmerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		logger.Warn("ID inválido ao buscar fazendeiro: %v", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	farmer, err := h.service.GetByID(uint(id))
	if err != nil {
		logger.Error("Erro ao buscar fazendeiro: %v", err)
		http.Error(w, "Erro ao buscar fazendeiro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farmer)
}

func (h *FarmerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters from query string
	params := models.PaginationParams{
		Page:  1,
		Limit: 10,
	}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			params.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			params.Limit = limit
		}
	}

	// Get paginated results from service
	result, err := h.service.GetAll(params)
	if err != nil {
		logger.Error("Erro ao buscar todos os fazendeiros: %v", err)
		http.Error(w, "Erro ao buscar fazendeiros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
