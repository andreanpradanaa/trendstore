package category

import (
	"errors"
	"testing"

	"github.com/andreanpradanaa/trendstore/internal/application/category/dto"
	"github.com/andreanpradanaa/trendstore/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Create(t *testing.T) {
	testCase := []struct {
		name      string
		mockSetup func(mockRepo *mocks.CategoryRepository)
		request   *dto.CreateCategoryRequest
		wantErr   bool
	}{
		{
			name: "Success - Create Category",
			mockSetup: func(mockRepo *mocks.CategoryRepository) {
				mockRepo.On("Create", mock.Anything).Return(nil)
			},
			request: &dto.CreateCategoryRequest{
				Name:        "fruits",
				Description: "Description",
			},
			wantErr: false,
		},
		{
			name: "Failed - Duplicate Name",
			mockSetup: func(mockRepo *mocks.CategoryRepository) {
				mockRepo.On("Create", mock.Anything).Return(errors.New("Duplicate category name"))
			},
			request: &dto.CreateCategoryRequest{
				Name:        "fruits",
				Description: "Description",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCase {
		mockRepo := mocks.NewCategoryRepository(t)
		tc.mockSetup(mockRepo)

		service := NewService(mockRepo)

		err := service.Create(tc.request)
		if tc.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
