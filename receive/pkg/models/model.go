package models

type Receive_Location struct {
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Capture_time string `json:"capture_time"`
}

type ReverseGeocode struct {
	Address []Addressparts `xml:"addressparts" json:"alamat"`
}

type Addressparts struct {
	Postcode       string `xml:"postcode" json:"kode_pos"`
	State_district string `xml:"state_district" json:"kota"`
}
