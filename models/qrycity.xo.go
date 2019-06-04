package models

// QryCity represents a row from 'public.qry_city'.
type QryCity struct {
	ID             NullInt64   `db:"id" json:"id" xml:"id"`                                           // id
	CountryID      NullString  `db:"country_id" json:"country_id" xml:"country_id"`                   // country_id
	DataSourceCode NullString  `db:"data_source_code" json:"data_source_code" xml:"data_source_code"` // data_source_code
	LoadDate       NullTime    `db:"load_date" json:"load_date" xml:"load_date"`                      // load_date
	CityIataCode   NullString  `db:"city_iata_code" json:"city_iata_code" xml:"city_iata_code"`       // city_iata_code
	CityCrtCode    NullString  `db:"city_crt_code" json:"city_crt_code" xml:"city_crt_code"`          // city_crt_code
	CityCladrCode  NullString  `db:"city_cladr_code" json:"city_cladr_code" xml:"city_cladr_code"`    // city_cladr_code
	CityNameRus    NullString  `db:"city_name_rus" json:"city_name_rus" xml:"city_name_rus"`          // city_name_rus
	CityNameEng    NullString  `db:"city_name_eng" json:"city_name_eng" xml:"city_name_eng"`          // city_name_eng
	Timezone       NullInt64   `db:"timezone" json:"timezone" xml:"timezone"`                         // timezone
	Longitude      NullFloat64 `db:"longitude" json:"longitude" xml:"longitude"`                      // longitude
	Latitude       NullFloat64 `db:"latitude" json:"latitude" xml:"latitude"`                         // latitude
	RegionName     NullString  `db:"region_name" json:"region_name" xml:"region_name"`                // region_name
	CityHashKey    NullInt64   `db:"city_hash_key" json:"city_hash_key" xml:"city_hash_key"`          // city_hash_key
	RowRank        NullInt64   `db:"row_rank" json:"row_rank" xml:"row_rank"`                         // row_rank
	CountryIsoCode NullString  `db:"country_iso_code" json:"country_iso_code" xml:"country_iso_code"` // country_iso_code
}
