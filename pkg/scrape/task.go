package scrape

	
type LebonCoinGeoZone struct {
    childs []LebonCoinGeoZone
    name string
}

var geoZones = &LebonCoinGeoZone{
	name: "france",
	childs : []LebonCoinGeoZone{
		{
			name: "rhone_alpes",
			childs: []LebonCoinGeoZone{
				{
					name: "ain",
					childs: []LebonCoinGeoZone{},
				},
				{
					name: "rhone",
					childs: []LebonCoinGeoZone{},
				},
			},
		},
		{
			name: "provence_alpes_cote_d_azur",
			childs: []LebonCoinGeoZone{},
		},
	},
}
