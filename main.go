package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dv_api/models"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/tkanos/gonfig"
)

func main() {
	fmt.Println("Service started")
	// Handle Subsequent requests
	handleRequests()
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/countries", allCountries).Methods("GET")
	myRouter.HandleFunc("/cities", allCities).Methods("GET")
	myRouter.HandleFunc("/airports", allAirports).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func allAirports(w http.ResponseWriter, r *http.Request) {
	source, sourceOk := checkSource(w, r)
	if sourceOk {

		db, err := connectToDB()
		if err != nil {
			panic("failed to connect database")
		}

		defer db.Close()
		airports := []models.QryAirport{}
		err = db.Select(&airports, "SELECT * FROM public.qry_airport WHERE data_source_code = $1", source)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(airports)
	}

}

func allCities(w http.ResponseWriter, r *http.Request) {
	source, sourceOk := checkSource(w, r)
	if sourceOk {

		db, err := connectToDB()
		if err != nil {
			panic("failed to connect database")
		}

		defer db.Close()
		cities := []models.QryCity{}
		err = db.Select(&cities, "SELECT * FROM public.qry_city WHERE data_source_code = $1", source)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(cities)
	}

}

func allCountries(w http.ResponseWriter, r *http.Request) {
	source, sourceOk := checkSource(w, r)
	if sourceOk {

		db, err := connectToDB()

		if err != nil {
			panic("failed to connect database")
		}

		defer db.Close()
		countries := []models.QryCountry{}
		err = db.Select(&countries, "SELECT * FROM public.qry_country WHERE data_source_code = $1", source)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(countries)

	}
}

func checkSource(w http.ResponseWriter, r *http.Request) (string, bool) {
	source, sourceOk := r.URL.Query()["source"]
	if !sourceOk || len(source[0]) < 1 {
		w.Write([]byte("Url Param 'source' is missing"))
		return "", false
	}
	return source[0], sourceOk
}

func connectToDB() (*sqlx.DB, error) {
	conf := databaseConfig{}
	confError := gonfig.GetConf("dbconfig.cfg", &conf)
	if confError != nil {
		fmt.Println(confError)
	}

	connString := "host=" + conf.Host + " port=" + conf.Port + " user=" + conf.User + " dbname=" + conf.Database + " password=" + conf.Password
	db, err := sqlx.Connect("postgres", connString)
	return db, err
}
