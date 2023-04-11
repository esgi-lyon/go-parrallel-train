package main

import (
	//"bufio"
	"encoding/csv"
	//"io"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"

	//"github.com/loic-roux-404/go-parrallel-train/pkg/lobstrio"
	//"github.com/loic-roux-404/go-parrallel-train/pkg/scrape"
)

type Asset struct {
	Id    string
	InputUrl     string
	Object string
	AnnonceId string
	ApiKey string
	Adtype string
	Area string
	CategoryName string
	Currency string
	City string
	ChargesIncluded string
	Capacity string
	Description string
	Department string
	DpeString string
	DpeInt string
	Dpe string
	District string
	DetailedTime string
	ExpirationDate string
	FirstPublicationDate string
	Floor string
	Furnished string
	GesString string
	GesInt string
	Ges string
	HasPhone string
	HasSwimmingPool string
	HasOnlineShop string
	IsExclusive string
	IsActive string
	IsDeactivated string
	IsDetailed string
	Lat string
	Lng string
	LastPublicationDate string
	Mail string
	MoreDetails string
	NoSalesmen string
	OwnerType string
	OwnerName string
	OnlineShopUrl string
	OwnerStoreId string
	OwnerSiren string
	Price string
	Picture string
	Pictures string
	PostalCode string
	Phone string
	PhoneFromUser string
	Region string
	RoomCount string
	RealEstateType string
	SleepingroomCount string
	SquareMetterPrice string
	StatusCode string
	Title string
	Url string
	Urgent string
	CollectedAt string
	UserId string
	Pieces string
	Surface string
	Honoraires string
	Reference string
	TypeDeBien string
	ClasseEnergie string
	NombreDeChambres string
	SurfaceDuTerrain string
	Ascenseur string
	EtageDuBien string
	PlacesDeParking string
}

func check(err error) {
    if err != nil {
        log.Panic(err)
    }
}

func main()  {
	for _, csv_file := range []string{
		"./data/Leboncoin Listings Search Export_20230418_1016.csv",
		"./data/Leboncoin Listings Search Export_20230411_1016.csv",
		"./data/Leboncoin Listings Search Export_20230425_1016.csv",
	} {
		dat, err := os.Open(csv_file)
    	check(err)

		reader := csv.NewReader(dat)

		var data Asset

		config := &mapstructure.DecoderConfig{
			ErrorUnused: true,
			Result:      &data,
		}

		decoder, err := mapstructure.NewDecoder(config)
		if err != nil {
			log.Panic(err)
		}

		if err := decoder.Decode(reader); err != nil {
			log.Panic(err)
		}

		log.Println(data)
	}
}
