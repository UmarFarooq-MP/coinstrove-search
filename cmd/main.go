package main

import (
	"coinstrove-search/internal/core/domain/dto"
	"coinstrove-search/internal/core/ports"
	"coinstrove-search/internal/core/services/datasvc"
	"coinstrove-search/internal/infrastructures/http"
	"coinstrove-search/repositories/mongo"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func readFromQueue(dataSVCHandler *http.DataSVCHandler) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed Initializing Broker Connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"prices", // name
		false,    // durable
		true,     // delete when unused
		false,    // exclusive
		true,     // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue with error : %v", err)
	}

	msgs, err := ch.Consume(
		"prices",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			message := dto.Message{}
			if err = json.Unmarshal(d.Body, &message); err != nil {
				log.Printf("Unable to Unmarshal the message from que, message = \n %s \n", d.Body)
				continue
			}
			dataSVCHandler.UpdateCoinInfo(message)
		}
	}()
	<-forever
}

func main() {
	dataSVCRepo, err := mongo.NewMongoRepository("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.10.0",
		"coinstrove", 2)
	if err != nil {
		log.Fatalf("Error while creating MongoRepository with message : %v", err)
	}

	dataService := []ports.DataSVC{
		datasvc.NewDataService(dataSVCRepo),
	}
	dataSVCHandler := http.NewDataSVCHandler(dataService)
	dataSVCRouter := http.SetupRouter(dataSVCHandler)
	go readFromQueue(dataSVCHandler)

	if err = dataSVCRouter.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server with error message %v", err)
	}
}
