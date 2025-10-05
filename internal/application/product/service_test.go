package product

import (
	"errors"
	"testing"

	"github.com/andreanpradanaa/trendstore/internal/application/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/domain/product"
	"github.com/andreanpradanaa/trendstore/internal/shared"
	"github.com/andreanpradanaa/trendstore/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Create(t *testing.T) {
	// Fixed timestamp untuk konsistensi testing

	testCases := []struct {
		name        string
		request     *dto.ProductRequest
		mockSetup   func(*mocks.ProductRepository)
		expectedErr error
		expectError bool
	}{
		{
			name: "Success - Create product with valid data",
			request: &dto.ProductRequest{
				Name:        "Laptop Gaming",
				Description: "High performance gaming laptop",
				Price:       15000000,
				Stock:       10,
				CategoryID:  1,
			},
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("Create", mock.MatchedBy(func(p *product.Product) bool {
					return p.Name == "Laptop Gaming" &&
						p.Description == "High performance gaming laptop" &&
						p.Price == 15000000 &&
						p.Stock == 10 &&
						p.CategoryID == 1
				})).Return(nil).Once()
			},
			expectedErr: nil,
			expectError: false,
		},
		{
			name: "Success - Create product with zero price (free product)",
			request: &dto.ProductRequest{
				Name:        "Free E-book",
				Description: "Free programming e-book",
				Price:       0,
				Stock:       1000,
				CategoryID:  2,
			},
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("Create", mock.MatchedBy(func(p *product.Product) bool {
					return p.Price == 0 && p.Name == "Free E-book"
				})).Return(nil).Once()
			},
			expectedErr: nil,
			expectError: false,
		},
		{
			name: "Repository error - database failure",
			request: &dto.ProductRequest{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10000,
				Stock:       5,
				CategoryID:  1,
			},
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("Create", mock.Anything).
					Return(errors.New("database connection failed")).
					Once()
			},
			expectedErr: errors.New("database connection failed"),
			expectError: true,
		},
		{
			name: "Repository error - duplicate product",
			request: &dto.ProductRequest{
				Name:        "Duplicate Product",
				Description: "Should fail duplicate",
				Price:       20000,
				Stock:       15,
				CategoryID:  1,
			},
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("Create", mock.Anything).
					Return(errors.New("duplicate key error")).
					Once()
			},
			expectedErr: errors.New("duplicate key error"),
			expectError: true,
		},
		{
			name: "Edge case - minimal product data",
			request: &dto.ProductRequest{
				Name:       "Minimal Product",
				Price:      5000,
				Stock:      1,
				CategoryID: 1,
				// Description omitted
			},
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("Create", mock.MatchedBy(func(p *product.Product) bool {
					return p.Name == "Minimal Product" &&
						p.Description == "" &&
						p.Price == 5000 &&
						p.Stock == 1
				})).Return(nil).Once()
			},
			expectedErr: nil,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			mockRepo := new(mocks.ProductRepository)
			service := NewService(mockRepo)

			// Setup mock expectations
			tc.mockSetup(mockRepo)

			// Execute
			err := service.Create(tc.request)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.expectedErr != nil {
					assert.Equal(t, tc.expectedErr.Error(), err.Error())
				}
			} else {
				assert.NoError(t, err)
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetByID(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocks.ProductRepository)
		id        int64
		want      *dto.ProductResponse
		wantErr   bool
	}{
		{
			name: "Success - Get Product By ID",
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("GetByID", mock.Anything).Return(&product.Product{
					ID:          1,
					Name:        "NIKE",
					Description: "Sepatu",
					Price:       500,
					Stock:       10,
					CategoryID:  3,
				}, nil).Once()
			},
			id: 1,
			want: &dto.ProductResponse{
				ID:          1,
				Name:        "NIKE",
				Description: "Sepatu",
				Price:       500,
				Stock:       10,
				CategoryID:  3,
				Slug:        shared.GenerateSlug("NIKE"),
			},
			wantErr: false,
		},
		{
			name: "Error Repository - Get Product By ID",
			mockSetup: func(mockRepo *mocks.ProductRepository) {
				mockRepo.On("GetByID", mock.Anything).Return(nil, errors.New(mock.Anything)).Once()
			},
			id:      1,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockRepo := new(mocks.ProductRepository)
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			// Create service dengan mock
			service := NewService(mockRepo)

			// Execute method
			got, err := service.GetByID(tt.id)

			// Assertions
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_Update(t *testing.T) {
	tests := []struct {
		name        string
		productRepo *mocks.ProductRepository
		args        *dto.ProductUpdateRequest
		wantErr     bool
	}{
		{
			name: "Success - Update Product",
			productRepo: func() *mocks.ProductRepository {
				mockRepo := new(mocks.ProductRepository)
				mockRepo.On("GetByID", int64(1)).Return(&product.Product{
					ID:          1,
					Name:        "NIKE",
					Description: "Sepatu",
					Price:       500,
					Stock:       10,
					CategoryID:  3,
				}, nil).Once()
				mockRepo.On("Update", mock.Anything).Return(nil).Once()
				return mockRepo
			}(),
			args: &dto.ProductUpdateRequest{
				ID:          1,
				Name:        "ADIDAS",
				Description: "Sepatu Baru",
				Price:       600,
				Stock:       15,
				CategoryID:  4,
			},
			wantErr: false,
		},
		{
			name: "Error - Product Not Found",
			productRepo: func() *mocks.ProductRepository {
				mockRepo := new(mocks.ProductRepository)
				mockRepo.On("GetByID", int64(2)).Return(nil, errors.New("product not found")).Once()
				return mockRepo
			}(),
			args: &dto.ProductUpdateRequest{
				ID:          2,
				Name:        "ADIDAS",
				Description: "Sepatu Baru",
				Price:       600,
				Stock:       15,
				CategoryID:  4,
			},
			wantErr: true,
		},
		{
			name: "Error - Update Product Fails",
			productRepo: func() *mocks.ProductRepository {
				mockRepo := new(mocks.ProductRepository)
				mockRepo.On("GetByID", int64(3)).Return(&product.Product{
					ID:          3,
					Name:        "PUMA",
					Description: "Sepatu Lama",
					Price:       400,
					Stock:       5,
					CategoryID:  2,
				}, nil).Once()
				mockRepo.On("Update", mock.Anything).Return(errors.New("update failed")).Once()
				return mockRepo
			}(),
			args: &dto.ProductUpdateRequest{
				ID:          3,
				Name:        "PUMA Updated",
				Description: "Sepatu Lama Updated",
				Price:       450,
				Stock:       8,
				CategoryID:  2,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.productRepo)
			err := service.UpdateProduct(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.productRepo.AssertExpectations(t)
		})
	}
}

func TestService_Delete(t *testing.T) {
	tests := []struct {
		name        string
		productRepo *mocks.ProductRepository
		id          int64
		wantErr     bool
	}{
		{
			name: "Success - Delete Product",
			productRepo: func() *mocks.ProductRepository {
				mockRepo := new(mocks.ProductRepository)
				mockRepo.On("GetByID", int64(1)).Return(&product.Product{}, nil).Once()
				mockRepo.On("Delete", int64(1)).Return(nil).Once()
				return mockRepo
			}(),
			id:      1,
			wantErr: false,
		},
		{
			name: "Error - Delete Non-Existent Product",
			productRepo: func() *mocks.ProductRepository {
				mockRepo := new(mocks.ProductRepository)
				mockRepo.On("GetByID", int64(2)).Return(&product.Product{}, errors.New("product not found")).Once()
				return mockRepo
			}(),
			id:      2,
			wantErr: true,
		},
		{
			name: "Error - Repository Failure on Delete",
			productRepo: func() *mocks.ProductRepository {
				mockRepo := new(mocks.ProductRepository)
				mockRepo.On("GetByID", int64(3)).Return(&product.Product{}, nil).Once()
				mockRepo.On("Delete", int64(3)).Return(errors.New("database error")).Once()
				return mockRepo
			}(),
			id:      3,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService(tt.productRepo)
			err := service.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.productRepo.AssertExpectations(t)
		})
	}
}
