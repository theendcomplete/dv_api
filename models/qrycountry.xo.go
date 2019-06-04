package models

// QryCountry represents a row from 'public.qry_country'.
type QryCountry struct {
	ID             NullString `db:"id" json:"id" xml:"id"`                                           // id
	DataSourceCode NullString `db:"data_source_code" json:"data_source_code" xml:"data_source_code"` // data_source_code
	LoadDate       NullTime   `db:"load_date" json:"load_date" xml:"load_date"`                      // load_date
	CountryIsoCode NullString `db:"country_iso_code" json:"country_iso_code" xml:"country_iso_code"` // country_iso_code
	CountryCrtCode NullString `db:"country_crt_code" json:"country_crt_code" xml:"country_crt_code"` // country_crt_code
	CountryNameRus NullString `db:"country_name_rus" json:"country_name_rus" xml:"country_name_rus"` // country_name_rus
	CountryNameEng NullString `db:"country_name_eng" json:"country_name_eng" xml:"country_name_eng"` // country_name_eng
	RegionNameRus  NullString `db:"region_name_rus" json:"region_name_rus" xml:"region_name_rus"`    // region_name_rus
	RegionNameEng  NullString `db:"region_name_eng" json:"region_name_eng" xml:"region_name_eng"`    // region_name_eng
	RowRank        NullInt64  `db:"row_rank" json:"row_rank" xml:"row_rank"`                         // row_rank
}
