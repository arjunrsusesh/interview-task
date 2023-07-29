package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"sync"
	"task/cfg"
	"task/internal/app/model"
	"task/pkg/db"
)

func InitGPG(c cfg.DBConfig) {
	log.Println("initializing database")
	err := db.DBConnection(c.User, c.DBName, c.Password, c.Host, c.Port)
	if err != nil {
		log.Fatal("error establishing database connection ", err)
	}
	db := db.GetDBConnection()
	query := "SELECT offer_id, client_id, country, image, image_width, image_height, text_locale, validity_text_locale, position, valid_from, show_from, valid_to, flag, page_count, store_url, store_url_title, offer_home FROM offer_table"

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return
	}
	defer rows.Close()

	offerChannel := make(chan model.OfferCompany)

	var wg sync.WaitGroup
	maxGoroutines := determineMaxGoroutines()

	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go fetchDataFromDatabase(db, i, maxGoroutines, offerChannel, &wg)
	}
	var offers []model.OfferCompany
	go func() {
		for offer := range offerChannel {
			offers = append(offers, offer)
		}
	}()

	wg.Wait()

	fmt.Println("Data retrieval and RAM storage are complete!")
}

func determineMaxGoroutines() int {
	numCPU := runtime.NumCPU()
	return numCPU
}

func fetchDataFromDatabase(db *sql.DB, startIndex, maxGoroutines int, offerChannel chan<- model.OfferCompany, wg *sync.WaitGroup) {
	defer wg.Done()
	query := fmt.Sprintf("SELECT offer_id, client_id, country, image, image_width, image_height, text_locale, validity_text_locale, position, valid_from, show_from, valid_to, flag, page_count, store_url, store_url_title, offer_home FROM offer_table WHERE MOD(offer_id, %d) = %d", maxGoroutines, startIndex)

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var offer model.OfferCompany
		err := rows.Scan(&offer.OfferID, &offer.ClientID, &offer.Country, &offer.Image, &offer.ImageWidth, &offer.ImageHeight, &offer.TextLocale, &offer.ValidityTextLocale, &offer.Position, &offer.ValidFrom, &offer.ShowFrom, &offer.ValidTo, &offer.Flag, &offer.PageCount, &offer.StoreURL, &offer.StoreURLTitle, &offer.OfferHome)
		if err != nil {
			log.Printf("Error scanning data: %v", err)
			continue
		}
		offerChannel <- offer
	}
}
