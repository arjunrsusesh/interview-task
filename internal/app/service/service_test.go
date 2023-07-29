package service

import (
	"errors"
	"task/internal/app/dto"
	"task/internal/app/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockEmployeeRepo struct{}

func (m *MockEmployeeRepo) GetDetailsByCountry(country string) ([]model.OfferCompany, error) {
	return []model.OfferCompany{
		{
			OfferID:       1,
			ClientID:      101,
			Country:       "TestCountry",
			Image:         "test_image.jpg",
			ImageWidth:    100,
			ImageHeight:   200,
			TextLocale:    "en_US",
			ValidFrom:     time.Now(),
			ShowFrom:      time.Now(),
			ValidTo:       time.Now().AddDate(0, 0, 7),
			Flag:          1,
			PageCount:     5,
			StoreURL:      "https://example.com/store",
			StoreURLTitle: "Test Store",
			OfferHome:     1,
		},
	}, nil
}

func TestStatusServiceImpl_GetDetailsByCountry(t *testing.T) {
	mockEmployeeRepo := &MockEmployeeRepo{}

	statusService := NewStatusService(mockEmployeeRepo)

	country := "TestCountry"
	expectedResponse := []dto.OfferCompanyResponse{
		{
			OfferID:       1,
			ClientID:      101,
			Country:       "TestCountry",
			Image:         "test_image.jpg",
			ImageWidth:    100,
			ImageHeight:   200,
			TextLocale:    "en_US",
			ValidFrom:     time.Now(),
			ShowFrom:      time.Now(),
			ValidTo:       time.Now().AddDate(0, 0, 7),
			Flag:          1,
			PageCount:     5,
			StoreURL:      "https://example.com/store",
			StoreURLTitle: "Test Store",
			OfferHome:     1,
		},
	}

	response, err := statusService.GetDetailsByCountry(country)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)

	country = "NonExistentCountry"
	expectedError := errors.New("failed to get values")

	response, err = statusService.GetDetailsByCountry(country)

	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, response)
}
