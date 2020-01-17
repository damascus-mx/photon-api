package service

import (
	"fmt"
	"log"

	env "github.com/damascus-mx/photon-api/authentication/common/config"
	"github.com/streadway/amqp"
)

// InitMQBroker Start MQ Broker
func InitMQBroker() *amqp.Connection {
	conn, err := amqp.Dial(env.MQ_CONNECTION)
	if err != nil {
		panic(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf(env.FailedService, "MQ Broker"))
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"auth_events", // name of the exchange
		"direct",      // type
		false,         // durable
		false,         // delete when complete
		false,         // internal
		false,         // noWait
		nil,           // arguments
	)

	log.Printf(env.ServiceConnected, "MQ Broker")
	return conn
}
