package config

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang_api_queue/receive/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SendToApi(body_queue string) {
	var rec_loc models.Receive_Location
	var reverse models.ReverseGeocode
	json.Unmarshal([]byte(body_queue), &rec_loc)
	apiUrl := "https://nominatim.openstreetmap.org/reverse?lat=" + rec_loc.Latitude + "&lon=" + rec_loc.Longitude
	req, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	xml.Unmarshal(body, &reverse)
	var jsonReverse models.ReverseGeocode
	jsonData, _ := json.Marshal(reverse)
	json.Unmarshal([]byte(jsonData), &jsonReverse)
	db := connect()
	defer db.Close()
	_, errs := db.Exec("insert into geolocate(longitude,latitude,postcode,state_district,capture_time)values(?,?,?,?,?)",
		rec_loc.Longitude,
		rec_loc.Latitude,
		jsonReverse.Address[0].Postcode,
		jsonReverse.Address[0].State_district,
		rec_loc.Capture_time,
	)
	if errs != nil {
		fmt.Println(err)
	}
}
