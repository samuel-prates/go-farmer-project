// internal/api/handlers/dashboard_handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samuel-prates/farm-project/backend/pkg/logger"
)

type DashboardHandler struct {
	service DashboardServiceInterface
}

func NewDashboardHandler(service DashboardServiceInterface) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetDashboardData(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetDashboardData()
	if err != nil {
		logger.Error("Erro ao buscar dados do dashboard: %v", err)
		http.Error(w, "Erro ao buscar dados do dashboard: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *DashboardHandler) GetFarmsByState(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetFarmsByState()
	if err != nil {
		logger.Error("Erro ao buscar fazendas por estado: %v", err)
		http.Error(w, "Erro ao buscar fazendas por estado: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *DashboardHandler) GetHarvestTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetHarvestTypes()
	if err != nil {
		logger.Error("Erro ao buscar tipos de cultivo: %v", err)
		http.Error(w, "Erro ao buscar tipos de cultivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *DashboardHandler) GetAreaDistribution(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAreaDistribution()
	if err != nil {
		logger.Error("Erro ao buscar distribuição de áreas: %v", err)
		http.Error(w, "Erro ao buscar distribuição de áreas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
