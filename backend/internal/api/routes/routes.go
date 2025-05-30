// internal/api/routes/routes.go
package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	routeHandlers "github.com/samuel-prates/farm-project/backend/internal/api/handlers"
)

func SetupRoutes(
	farmerHandler *routeHandlers.FarmerHandler,
	dashboardHandler *routeHandlers.DashboardHandler,
) http.Handler {
	r := mux.NewRouter()

	// Rotas para Fazendeiros
	r.HandleFunc("/api/farmers", farmerHandler.Create).Methods("POST")
	r.HandleFunc("/api/farmers/{id}", farmerHandler.Update).Methods("PUT")
	r.HandleFunc("/api/farmers/{id}", farmerHandler.Delete).Methods("DELETE")
	r.HandleFunc("/api/farmers/{id}", farmerHandler.GetByID).Methods("GET")
	r.HandleFunc("/api/farmers", farmerHandler.GetAll).Methods("GET")

	// Rotas para Dashboard
	r.HandleFunc("/api/dashboard", dashboardHandler.GetDashboardData).Methods("GET")
	r.HandleFunc("/api/dashboard/farm-states", dashboardHandler.GetFarmsByState).Methods("GET")
	r.HandleFunc("/api/dashboard/harvest-cultures", dashboardHandler.GetHarvestTypes).Methods("GET")
	r.HandleFunc("/api/dashboard/areas", dashboardHandler.GetAreaDistribution).Methods("GET")

	// Add CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*", "http://localhost:*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return corsMiddleware(r)
}
