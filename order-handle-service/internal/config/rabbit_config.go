package  config

import (
    amqp "github.com/rabbitmq/amqp091-go"
    "log"
)

var ConnRabbitMq *amqp.Connection

func ConfigRabbitMq() {
    conn, err :=  amqp.Dial(Envs.RabbitBrokerUrl)

    if err !=  nil {
        log.Panicf("Error on connect in RabbitMQ: %s \n", err.Error())
    }

    ConnRabbitMq = conn
}