package repo

import (
	"database/sql"
	"log"
	"task/internal/app/model"
	"task/pkg/db"
)

type CompanyRepo interface {
	GetDetailsByCountry(string) ([]model.OfferCompany, error)
}

type companyRepoImpl struct {
	db *sql.DB
}

func NewCompanyRepo(db *sql.DB) CompanyRepo {
	return &companyRepoImpl{
		db: db,
	}
}

func (r *companyRepoImpl) GetDetailsByCountry(country string) ([]model.OfferCompany, error) {
	db := db.GetDBConnection()
	var companees []model.OfferCompany
	query := "SELECT * FROM offer_company WHERE country = $1"
	rows, err := db.Query(query, country)
	if err != nil {
		log.Panicln("failed to connect to database in repo")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var company model.OfferCompany
		if err := rows.Scan(&company.OfferID, &company.ClientID, &company.Country, &company.Image, &company.ImageWidth, &company.ImageHeight, &company.TextLocale, &company.ValidityTextLocale, &company.Position, &company.ValidFrom, &company.ShowFrom, &company.ValidTo, &company.Flag, &company.PageCount, &company.StoreURL, &company.StoreURLTitle, &company.OfferHome); err != nil {
			log.Println("failed to scan values")
			continue
		}
		companees = append(companees, company)
	}

	return companees, nil
}
