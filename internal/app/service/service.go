package service

import (
	"log"
	"task/internal/app/dto"
	"task/internal/app/model"
	"task/internal/app/repo"
)

type StatusService interface {
	GetDetailsByCountry(string) ([]dto.OfferCompanyResponse, error)
}

type statusServiceImpl struct {
	CompanyRepo repo.CompanyRepo
}

func NewStatusService(companyRepo repo.CompanyRepo) StatusService {
	return &statusServiceImpl{
		CompanyRepo: companyRepo,
	}
}
func (s *statusServiceImpl) GetDetailsByCountry(country string) ([]dto.OfferCompanyResponse, error) {
	result, err := s.CompanyRepo.GetDetailsByCountry(country)
	if err != nil {
		log.Println("failed to get values")
		return nil, err
	}
	var responseDTOs []dto.OfferCompanyResponse
	for _, modelCompany := range result {
		dtoCompany := convertToDTO(modelCompany)
		responseDTOs = append(responseDTOs, dtoCompany)
	}

	return responseDTOs, nil
}
func convertToDTO(modelOfferCompany model.OfferCompany) dto.OfferCompanyResponse {
	return dto.OfferCompanyResponse{
		OfferID:            uint32(modelOfferCompany.OfferID),
		ClientID:           uint32(modelOfferCompany.ClientID),
		Country:            modelOfferCompany.Country,
		Image:              modelOfferCompany.Image,
		ImageWidth:         int32(modelOfferCompany.ImageWidth),
		ImageHeight:        int32(modelOfferCompany.ImageHeight),
		TextLocale:         modelOfferCompany.TextLocale,
		ValidityTextLocale: modelOfferCompany.ValidityTextLocale,
		Position:           int32(modelOfferCompany.Position),
		ValidFrom:          modelOfferCompany.ValidFrom,
		ShowFrom:           modelOfferCompany.ShowFrom,
		ValidTo:            modelOfferCompany.ValidTo,
		Flag:               uint32(modelOfferCompany.Flag),
		PageCount:          uint32(modelOfferCompany.PageCount),
		StoreURL:           modelOfferCompany.StoreURL,
		StoreURLTitle:      modelOfferCompany.StoreURLTitle,
		OfferHome:          int32(modelOfferCompany.OfferHome),
	}
}
