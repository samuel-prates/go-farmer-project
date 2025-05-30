// internal/api/handlers/dashboard_handler_test.go
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samuel-prates/farm-project/backend/internal/models"
)

// MockDashboardService is a mock implementation of the DashboardServiceInterface
type MockDashboardService struct {
	GetDashboardDataFunc    func() (*DashboardData, error)
	GetFarmsByStateFunc     func() ([]models.StateCount, error)
	GetHarvestTypesFunc     func() ([]models.HarvestCultureCount, error)
	GetAreaDistributionFunc func() (*AreaDistribution, error)
}

func (m *MockDashboardService) GetDashboardData() (*DashboardData, error) {
	return m.GetDashboardDataFunc()
}

func (m *MockDashboardService) GetFarmsByState() ([]models.StateCount, error) {
	return m.GetFarmsByStateFunc()
}

func (m *MockDashboardService) GetHarvestTypes() ([]models.HarvestCultureCount, error) {
	return m.GetHarvestTypesFunc()
}

func (m *MockDashboardService) GetAreaDistribution() (*AreaDistribution, error) {
	return m.GetAreaDistributionFunc()
}

func TestDashboardHandler_GetDashboardData(t *testing.T) {
	// Test cases
	tests := []struct {
		name                   string
		mockGetDashboardDataFunc func() (*DashboardData, error)
		expectedStatus         int
	}{
		{
			name: "Success",
 		mockGetDashboardDataFunc: func() (*DashboardData, error) {
 			return &DashboardData{
 				TotalFarms: 10,
 				TotalArea:  1000.5,
 			}, nil
 		},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service Error",
 		mockGetDashboardDataFunc: func() (*DashboardData, error) {
 			return nil, errors.New("service error")
 		},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockDashboardService{
				GetDashboardDataFunc: tt.mockGetDashboardDataFunc,
			}
			handler := NewDashboardHandler(mockService)

			// Create request
			req, err := http.NewRequest("GET", "/api/dashboard", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetDashboardData(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// For successful responses, verify the response body
			if tt.expectedStatus == http.StatusOK {
				var response DashboardData
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}

				expectedData, _ := tt.mockGetDashboardDataFunc()
				if response.TotalFarms != expectedData.TotalFarms || response.TotalArea != expectedData.TotalArea {
					t.Errorf("Handler returned unexpected body: got %+v want %+v", response, expectedData)
				}
			}
		})
	}
}

func TestDashboardHandler_GetFarmsByState(t *testing.T) {
	// Test cases
	tests := []struct {
		name                  string
		mockGetFarmsByStateFunc func() ([]models.StateCount, error)
		expectedStatus        int
	}{
		{
			name: "Success",
			mockGetFarmsByStateFunc: func() ([]models.StateCount, error) {
				return []models.StateCount{
					{State: "SP", Count: 5},
					{State: "MG", Count: 3},
				}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Empty List",
			mockGetFarmsByStateFunc: func() ([]models.StateCount, error) {
				return []models.StateCount{}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service Error",
			mockGetFarmsByStateFunc: func() ([]models.StateCount, error) {
				return nil, errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockDashboardService{
				GetFarmsByStateFunc: tt.mockGetFarmsByStateFunc,
			}
			handler := NewDashboardHandler(mockService)

			// Create request
			req, err := http.NewRequest("GET", "/api/dashboard/farms-by-state", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetFarmsByState(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestDashboardHandler_GetHarvestTypes(t *testing.T) {
	// Test cases
	tests := []struct {
		name                  string
		mockGetHarvestTypesFunc func() ([]models.HarvestCultureCount, error)
		expectedStatus        int
	}{
		{
			name: "Success",
			mockGetHarvestTypesFunc: func() ([]models.HarvestCultureCount, error) {
				return []models.HarvestCultureCount{
					{Culture: "Soja", Count: 7},
					{Culture: "Milho", Count: 4},
				}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Empty List",
			mockGetHarvestTypesFunc: func() ([]models.HarvestCultureCount, error) {
				return []models.HarvestCultureCount{}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service Error",
			mockGetHarvestTypesFunc: func() ([]models.HarvestCultureCount, error) {
				return nil, errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockDashboardService{
				GetHarvestTypesFunc: tt.mockGetHarvestTypesFunc,
			}
			handler := NewDashboardHandler(mockService)

			// Create request
			req, err := http.NewRequest("GET", "/api/dashboard/harvest-types", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetHarvestTypes(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestDashboardHandler_GetAreaDistribution(t *testing.T) {
	// Test cases
	tests := []struct {
		name                       string
		mockGetAreaDistributionFunc func() (*AreaDistribution, error)
		expectedStatus             int
	}{
		{
			name: "Success",
 		mockGetAreaDistributionFunc: func() (*AreaDistribution, error) {
 			return &AreaDistribution{
 				AgricultureArea: 750.5,
 				VegetationArea:  250.0,
 			}, nil
 		},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service Error",
 		mockGetAreaDistributionFunc: func() (*AreaDistribution, error) {
 			return nil, errors.New("service error")
 		},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockDashboardService{
				GetAreaDistributionFunc: tt.mockGetAreaDistributionFunc,
			}
			handler := NewDashboardHandler(mockService)

			// Create request
			req, err := http.NewRequest("GET", "/api/dashboard/area-distribution", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetAreaDistribution(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// For successful responses, verify the response body
			if tt.expectedStatus == http.StatusOK {
				var response AreaDistribution
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}

				expectedData, _ := tt.mockGetAreaDistributionFunc()
				if response.AgricultureArea != expectedData.AgricultureArea || response.VegetationArea != expectedData.VegetationArea {
					t.Errorf("Handler returned unexpected body: got %+v want %+v", response, expectedData)
				}
			}
		})
	}
}
