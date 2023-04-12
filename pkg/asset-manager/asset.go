package assetmanager

import "time"

type Asset struct {
	ID                   string      `json:"id"`
	Object               string      `json:"object"`
	Cluster              string      `json:"cluster"`
	Task                 string      `json:"task"`
	Dpe                  string      `json:"DPE"`
	DPEString            string      `json:"DPE_string"`
	Ges                  string      `json:"GES"`
	GESString            string      `json:"GES_string"`
	AdType               string      `json:"ad_type"`
	AnnonceID            string      `json:"annonce_id"`
	Area                 int         `json:"area"`
	AssociationTime      time.Time   `json:"association_time"`
	CategoryName         string      `json:"category_name"`
	City                 string      `json:"city"`
	Currency             string      `json:"currency"`
	Department           string      `json:"department"`
	Description          string      `json:"description"`
	Details              Details     `json:"details"`
	FirstPublicationDate time.Time   `json:"first_publication_date"`
	LastPublicationDate  time.Time   `json:"last_publication_date"`
	Lat                  string      `json:"lat"`
	Lng                  string      `json:"lng"`
	MoreDetails          MoreDetails `json:"more_details"`
	NoSalesmen           bool        `json:"no_salesmen"`
	OwnerName            string      `json:"owner_name"`
	OwnerSiren           string      `json:"owner_siren"`
	OwnerStoreID         string      `json:"owner_store_id"`
	OwnerType            string      `json:"owner_type"`
	Phone                string      `json:"phone"`
	Pictures             string      `json:"pictures"`
	PostalCode           string      `json:"postal_code"`
	Price                int         `json:"price"`
	RealEstateType       string      `json:"real_estate_type"`
	Region               string      `json:"region"`
	RoomCount            int         `json:"room_count"`
	ScrapingTime         time.Time   `json:"scraping_time"`
	Title                string      `json:"title"`
	Urgent               bool        `json:"urgent"`
	URL                  string      `json:"url"`
}
type Details struct {
	Ges              string `json:"GES"`
	PiCes            string `json:"Pièces"`
	Surface          string `json:"Surface"`
	Honoraires       string `json:"Honoraires"`
	RFRence          string `json:"Référence"`
	TypeDeBien       string `json:"Type de bien"`
	ClasseNergie     string `json:"Classe énergie"`
	NombreDeChambres string `json:"Nombre de chambres"`
}
type MoreDetails struct {
	Ges                    string `json:"ges"`
	Rooms                  string `json:"rooms"`
	Square                 string `json:"square"`
	Bedrooms               string `json:"bedrooms"`
	IsImport               string `json:"is_import"`
	CustomRef              string `json:"custom_ref"`
	LeaseType              string `json:"lease_type"`
	DistrictID             string `json:"district_id"`
	EnergyRate             string `json:"energy_rate"`
	FaiIncluded            string `json:"fai_included"`
	MandateType            string `json:"mandate_type"`
	ExternalAdID           string `json:"external_ad_id"`
	ImmoSellType           string `json:"immo_sell_type"`
	ProRatesLink           string `json:"pro_rates_link"`
	ActivitySector         string `json:"activity_sector"`
	DistrictTypeID         string `json:"district_type_id"`
	RealEstateType         string `json:"real_estate_type"`
	DistrictVisibility     string `json:"district_visibility"`
	DistrictResolutionType string `json:"district_resolution_type"`
}
