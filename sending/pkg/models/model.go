package models

type Location_Request struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Capture_time string `json:"capture_time"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Location_Request
}
