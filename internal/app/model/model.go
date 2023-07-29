package model

import "time"

type OfferCompany struct {
	OfferID            uint32    `json:"offer_id"`
	ClientID           uint32    `json:"client_id"`
	Country            string    `json:"country"`
	Image              string    `json:"image"`
	ImageWidth         int32     `json:"image_width"`
	ImageHeight        int32     `json:"image_height"`
	TextLocale         string    `json:"text_locale"`
	ValidityTextLocale string    `json:"validity_text_locale"`
	Position           int32     `json:"position"`
	ValidFrom          time.Time `json:"valid_from"`
	ShowFrom           time.Time `json:"show_from"`
	ValidTo            time.Time `json:"valid_to"`
	Flag               uint32    `json:"flag"`
	PageCount          uint32    `json:"page_count"`
	StoreURL           string    `json:"store_url"`
	StoreURLTitle      string    `json:"store_url_title"`
	OfferHome          int32     `json:"offer_home"`
}
