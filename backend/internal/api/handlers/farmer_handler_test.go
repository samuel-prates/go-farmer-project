// internal/api/handlers/farmer_handler_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/samuel-prates/farm-project/backend/internal/models"
)

// MockFarmerService is a mock implementation of the FarmerServiceInterface
type MockFarmerService struct {
	CreateFunc  func(farmer *models.Farmer) (*models.Farmer, error)
	UpdateFunc  func(farmer *models.Farmer) (*models.Farmer, error)
	DeleteFunc  func(id uint) error
	GetByIDFunc func(id uint) (*models.Farmer, error)
	GetAllFunc  func(params models.PaginationParams) (models.PaginatedResult, error)
}

func (m *MockFarmerService) Create(farmer *models.Farmer) (*models.Farmer, error) {
	return m.CreateFunc(farmer)
}

func (m *MockFarmerService) Update(farmer *models.Farmer) (*models.Farmer, error) {
	return m.UpdateFunc(farmer)
}

func (m *MockFarmerService) Delete(id uint) error {
	return m.DeleteFunc(id)
}

func (m *MockFarmerService) GetByID(id uint) (*models.Farmer, error) {
	return m.GetByIDFunc(id)
}

func (m *MockFarmerService) GetAll(params models.PaginationParams) (models.PaginatedResult, error) {
	return m.GetAllFunc(params)
}

func TestFarmerHandler_Create(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		requestBody    interface{}
		mockCreateFunc func(farmer *models.Farmer) (*models.Farmer, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			requestBody: models.Farmer{
				FarmerName:            "Test Farmer",
				FederalIdentification: "12345678901",
			},
			mockCreateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				farmer.ID = 1
				return farmer, nil
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:        "Invalid JSON",
			requestBody: "invalid json",
			mockCreateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Validation Error",
			requestBody: models.Farmer{
				FarmerName: "", // Missing required field
			},
			mockCreateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Service Error",
			requestBody: models.Farmer{
				FarmerName:            "Test Farmer",
				FederalIdentification: "12345678901",
			},
			mockCreateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockFarmerService{
				CreateFunc: tt.mockCreateFunc,
			}
			handler := NewFarmerHandler(mockService)

			// Create request
			var reqBody []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req, err := http.NewRequest("POST", "/api/farmers", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.Create(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestFarmerHandler_GetByID(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		farmerID       string
		mockGetByIDFunc func(id uint) (*models.Farmer, error)
		expectedStatus int
	}{
		{
			name:     "Success",
			farmerID: "1",
			mockGetByIDFunc: func(id uint) (*models.Farmer, error) {
				return &models.Farmer{
					ID:                    id,
					FarmerName:            "Test Farmer",
					FederalIdentification: "12345678901",
				}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:     "Invalid ID",
			farmerID: "invalid",
			mockGetByIDFunc: func(id uint) (*models.Farmer, error) {
				return nil, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "Not Found",
			farmerID: "999",
			mockGetByIDFunc: func(id uint) (*models.Farmer, error) {
				return nil, errors.New("farmer not found")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockFarmerService{
				GetByIDFunc: tt.mockGetByIDFunc,
			}
			handler := NewFarmerHandler(mockService)

			// Create request
			req, err := http.NewRequest("GET", "/api/farmers/"+tt.farmerID, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Add URL parameters to request
			vars := map[string]string{
				"id": tt.farmerID,
			}
			req = mux.SetURLVars(req, vars)

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetByID(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestFarmerHandler_GetAll(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		mockGetAllFunc func() ([]models.Farmer, error)
		expectedStatus int
	}{
		{
			name: "Success",
			mockGetAllFunc: func() ([]models.Farmer, error) {
				return []models.Farmer{
					{
						ID:                    1,
						FarmerName:            "Farmer 1",
						FederalIdentification: "12345678901",
					},
					{
						ID:                    2,
						FarmerName:            "Farmer 2",
						FederalIdentification: "10987654321",
					},
				}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Empty List",
			mockGetAllFunc: func() ([]models.Farmer, error) {
				return []models.Farmer{}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Service Error",
			mockGetAllFunc: func() ([]models.Farmer, error) {
				return nil, errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockFarmerService{
				GetAllFunc: tt.mockGetAllFunc,
			}
			handler := NewFarmerHandler(mockService)

			// Create request
			req, err := http.NewRequest("GET", "/api/farmers", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetAll(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestFarmerHandler_Update(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		farmerID       string
		requestBody    interface{}
		mockUpdateFunc func(farmer *models.Farmer) (*models.Farmer, error)
		expectedStatus int
	}{
		{
			name:     "Success",
			farmerID: "1",
			requestBody: models.Farmer{
				FarmerName:            "Updated Farmer",
				FederalIdentification: "12345678901",
			},
			mockUpdateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return farmer, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:     "Invalid ID",
			farmerID: "invalid",
			requestBody: models.Farmer{
				FarmerName:            "Updated Farmer",
				FederalIdentification: "12345678901",
			},
			mockUpdateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Invalid JSON",
			farmerID:    "1",
			requestBody: "invalid json",
			mockUpdateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "Validation Error",
			farmerID: "1",
			requestBody: models.Farmer{
				FarmerName: "", // Missing required field
			},
			mockUpdateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "Service Error",
			farmerID: "1",
			requestBody: models.Farmer{
				FarmerName:            "Updated Farmer",
				FederalIdentification: "12345678901",
			},
			mockUpdateFunc: func(farmer *models.Farmer) (*models.Farmer, error) {
				return nil, errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockFarmerService{
				UpdateFunc: tt.mockUpdateFunc,
			}
			handler := NewFarmerHandler(mockService)

			// Create request
			var reqBody []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req, err := http.NewRequest("PUT", "/api/farmers/"+tt.farmerID, bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Add URL parameters to request
			vars := map[string]string{
				"id": tt.farmerID,
			}
			req = mux.SetURLVars(req, vars)

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.Update(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestFarmerHandler_Delete(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		farmerID       string
		mockDeleteFunc func(id uint) error
		expectedStatus int
	}{
		{
			name:     "Success",
			farmerID: "1",
			mockDeleteFunc: func(id uint) error {
				return nil
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:     "Invalid ID",
			farmerID: "invalid",
			mockDeleteFunc: func(id uint) error {
				return nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:     "Service Error",
			farmerID: "1",
			mockDeleteFunc: func(id uint) error {
				return errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock service
			mockService := &MockFarmerService{
				DeleteFunc: tt.mockDeleteFunc,
			}
			handler := NewFarmerHandler(mockService)

			// Create request
			req, err := http.NewRequest("DELETE", "/api/farmers/"+tt.farmerID, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Add URL parameters to request
			vars := map[string]string{
				"id": tt.farmerID,
			}
			req = mux.SetURLVars(req, vars)

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.Delete(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
