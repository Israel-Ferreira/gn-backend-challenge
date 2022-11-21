package publishers

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/internal/config"
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/models"
)

func PublishAddress(address models.Address) {
	ch, err := config.ConnRabbitMq.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare("addresses", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	msg := models.ToAddressMsg(address)

	jsonMsg, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	err = ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         jsonMsg,
	})

	if err != nil {
		panic(err)
	}

}
