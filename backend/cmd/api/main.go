// cmd/api/main.go
package main

import (
	"net/http"
	"os"
	"time"

	"github.com/samuel-prates/farm-project/backend/internal/api/handlers"
	"github.com/samuel-prates/farm-project/backend/internal/api/routes"
	"github.com/samuel-prates/farm-project/backend/internal/repository"
	"github.com/samuel-prates/farm-project/backend/internal/services"
	"github.com/samuel-prates/farm-project/backend/pkg/config"
	"github.com/samuel-prates/farm-project/backend/pkg/database"
	"github.com/samuel-prates/farm-project/backend/pkg/logger"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Erro ao conectar ao banco de dados: %v", err)
	}

	// Inicializar repositórios
	farmerRepo := repository.NewFarmerRepository(db)
	farmRepo := repository.NewFarmRepository(db)
	harvestRepo := repository.NewHarvestRepository(db)

	// Inicializar serviços
	farmerService := services.NewFarmerService(farmerRepo)
	dashboardService := services.NewDashboardService(farmRepo, harvestRepo)

	// Inicializar handlers com adaptadores
	farmerHandler := handlers.NewFarmerHandler(handlers.NewFarmerServiceAdapter(farmerService))
	dashboardHandler := handlers.NewDashboardHandler(handlers.NewDashboardServiceAdapter(dashboardService))

	// Configurar rotas
	router := routes.SetupRoutes(farmerHandler, dashboardHandler)

	// Configurar servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Iniciar servidor
	logger.Info("Servidor iniciado na porta %s", port)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Erro ao iniciar servidor: %v", err)
	}
}
