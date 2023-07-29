package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDetailsByCountry(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer mockDB.Close()
	repo := NewCompanyRepo(mockDB)
	testCountry := "TestCountry"
	columns := []string{"offer_id", "client_id", "country", "image", "image_width", "image_height", "text_locale", "validity_text_locale", "position", "valid_from", "show_from", "valid_to", "flag", "page_count", "store_url", "store_url_title", "offer_home"}
	mock.ExpectQuery("SELECT .* FROM qburst WHERE country = .*").WithArgs(testCountry).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(307470, 1001, "US", "image1.jpg", 800, 600, "Special offer for you!", "Valid until end of the month", 1, "2023-07-01 00:00:00", "2023-07-01 00:00:00", "2023-07-31 23:59:59", 1, 5, "www.store1.com", "Store 1", true).
			AddRow(307471, 1002, "US", "image2.png", 1024, 768, "Exclusive discounts!", "Valid for one week", 2, "2023-07-12 00:00:00", "2023-07-12 00:00:00", "2023-07-19 23:59:59", 1, 3, "www.store2.co.uk", "Store 2", false))

	employees, err := repo.GetDetailsByCountry(testCountry)
	assert.NoError(t, err)
	assert.NotNil(t, employees)
	assert.Len(t, employees, 2)
	assert.Equal(t, "image1", employees[0].Image)
	assert.Equal(t, "image2", employees[1].Image)
	assert.NoError(t, mock.ExpectationsWereMet())
}
