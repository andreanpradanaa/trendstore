package usecase

import (
	"testing"

	"github.com/andreanpradanaa/trendstore/internal/product/dto"
	"github.com/andreanpradanaa/trendstore/internal/product/repository"
	"github.com/andreanpradanaa/trendstore/internal/product/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestProductUsecase_CreateProduct(t *testing.T) {
	type fields struct {
		ProductRepo repository.ProductRepository
	}

	type args struct {
		args *dto.ProductRequest
	}

	tests := []struct {
		name         string
		args         args
		mockBehavior func(m *mocks.ProductRepository)
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				args: &dto.ProductRequest{
					Name:        "Test Name",
					Description: "Desc",
					Price:       100,
					Stock:       1,
					CategoryID:  1,
				},
			},
			wantErr: false,
			mockBehavior: func(m *mocks.ProductRepository) {
				m.On("CreateProduct", mock.Anything).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProductRepo := mocks.NewProductRepository(t)

			if tt.mockBehavior != nil {
				tt.mockBehavior(mockProductRepo)
			}

			fields := fields{
				ProductRepo: mockProductRepo,
			}

			p := &productUsecase{
				ProductRepo: fields.ProductRepo,
			}
			if err := p.CreateProduct(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ProductUsecase.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
