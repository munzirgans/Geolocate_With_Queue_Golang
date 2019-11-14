package main

import (
	"fmt"
	"golang_api_queue/receive/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"geolocate",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			body_queue := string(d.Body)
			fmt.Println(fmt.Sprintf("Receive Message : %s", body_queue))
			config.SendToApi(body_queue)
		}
	}()
	fmt.Println("Waiting for messages")
	<-forever
}
