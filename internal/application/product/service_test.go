package product

import (
	"errors"
	"testing"

	"github.com/andreanpradanaa/trendstore/internal/application/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/domain/product"
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
