package config

import (
	"encoding/json"
	"fmt"
	"golang_api_queue/sending/pkg/models"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

func InsertLocation(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	var loc_req models.Location_Request
	var response models.Response
	json.NewDecoder(r.Body).Decode(&loc_req)
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer connection.Close()
	if err != nil {
		fmt.Println(err)
	}
	ch, err := connection.Channel()
	if err != nil {
		fmt.Println(err)
	}
	now_formatted := now.Format("2006-01-02 15:04:05")
	now_js := `{"capture_time":"` + now_formatted + `"}`
	json.Unmarshal([]byte(now_js), &loc_req)
	body, _ := json.Marshal(loc_req)
	fmt.Println(string(body))
	err = ch.Publish(
		"",
		"geolocate",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	response.Status = 1
	response.Message = "Your Data has been sent to the queue"
	response.Data = loc_req
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
