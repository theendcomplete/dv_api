package models

// QryAirport represents a row from 'public.qry_airport'.
type QryAirport struct {
	ID                  NullInt64   `db:"id" json:"id" xml:"id"`                                                             // id
	CityID              NullInt64   `db:"city_id" json:"city_id" xml:"city_id"`                                              // city_id
	AirportIataCode     NullString  `db:"airport_iata_code" json:"airport_iata_code" xml:"airport_iata_code"`                // airport_iata_code
	DataSourceCode      NullString  `db:"data_source_code" json:"data_source_code" xml:"data_source_code"`                   // data_source_code
	LoadDate            NullTime    `db:"load_date" json:"load_date" xml:"load_date"`                                        // load_date
	AirportIcaoCode     NullString  `db:"airport_icao_code" json:"airport_icao_code" xml:"airport_icao_code"`                // airport_icao_code
	AirportCrtCode      NullString  `db:"airport_crt_code" json:"airport_crt_code" xml:"airport_crt_code"`                   // airport_crt_code
	AirportShortNameEng NullString  `db:"airport_short_name_eng" json:"airport_short_name_eng" xml:"airport_short_name_eng"` // airport_short_name_eng
	AirportFullNameRus  NullString  `db:"airport_full_name_rus" json:"airport_full_name_rus" xml:"airport_full_name_rus"`    // airport_full_name_rus
	AirportFullNameEng  NullString  `db:"airport_full_name_eng" json:"airport_full_name_eng" xml:"airport_full_name_eng"`    // airport_full_name_eng
	EmailAddress        NullString  `db:"email_address" json:"email_address" xml:"email_address"`                            // email_address
	ContactPhoneNum     NullString  `db:"contact_phone_num" json:"contact_phone_num" xml:"contact_phone_num"`                // contact_phone_num
	Latitude            NullFloat64 `db:"latitude" json:"latitude" xml:"latitude"`                                           // latitude
	Longtitude          NullFloat64 `db:"longtitude" json:"longtitude" xml:"longtitude"`                                     // longtitude
	Website             NullString  `db:"website" json:"website" xml:"website"`                                              // website
	AirportHashKey      NullInt64   `db:"airport_hash_key" json:"airport_hash_key" xml:"airport_hash_key"`                   // airport_hash_key
	AirportShortNameRus NullString  `db:"airport_short_name_rus" json:"airport_short_name_rus" xml:"airport_short_name_rus"` // airport_short_name_rus
	RowRank             NullInt64   `db:"row_rank" json:"row_rank" xml:"row_rank"`                                           // row_rank
}
